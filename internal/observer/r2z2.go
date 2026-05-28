package observer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"
)

const (
	r2z2BaseURL  = "https://r2z2.zkillboard.com"
	httpTimeout  = 30 * time.Second
	seenCapacity = 10000
)

// seenSet is a bounded FIFO membership set used to deduplicate killmails by
// killmail_id. The same killmail_id can recur under a new sequence_id, so the
// poller checks membership before emitting. It is not safe for concurrent use;
// the poll loop runs in a single goroutine.
type seenSet struct {
	ids      map[int64]struct{}
	order    []int64
	capacity int
}

func newSeenSet(capacity int) *seenSet {
	return &seenSet{
		ids:      make(map[int64]struct{}, capacity),
		order:    make([]int64, 0, capacity),
		capacity: capacity,
	}
}

func (s *seenSet) seen(id int64) bool {
	_, ok := s.ids[id]
	return ok
}

// add records id, evicting the oldest entry once the set is at capacity.
func (s *seenSet) add(id int64) {
	if _, ok := s.ids[id]; ok {
		return
	}
	if len(s.ids) >= s.capacity && len(s.order) > 0 {
		oldest := s.order[0]
		s.order = s.order[1:]
		delete(s.ids, oldest)
	}
	s.ids[id] = struct{}{}
	s.order = append(s.order, id)
}

// r2z2 tails the zKillboard R2Z2 killmail feed. R2Z2 is a sequence of static
// JSON files behind a client-owned integer cursor: each killmail lives at
// /ephemeral/{sequence_id}.json, and /ephemeral/sequence.json reports the
// latest known sequence. The poller bootstraps from the latest sequence and
// tails forward, so killmails produced during downtime are missed.
type r2z2 struct {
	out       chan *Killmail
	client    *http.Client
	baseURL   string
	userAgent string
	active    time.Duration
	idle      time.Duration
	seen      *seenSet
	log       *zap.SugaredLogger
}

// newR2Z2 builds a poller. It does not start any goroutine; call run to begin
// tailing. userAgent must be non-blank: R2Z2 rejects blank User-Agent requests
// with 403.
func newR2Z2(out chan *Killmail, client *http.Client, baseURL, userAgent string, active, idle time.Duration, log *zap.SugaredLogger) *r2z2 {
	return &r2z2{
		out:       out,
		client:    client,
		baseURL:   baseURL,
		userAgent: userAgent,
		active:    active,
		idle:      idle,
		seen:      newSeenSet(seenCapacity),
		log:       log,
	}
}

// run bootstraps the starting sequence then tails the feed until ctx is
// cancelled. Bootstrap is retried on error (honoring cancellation) because the
// poller cannot advance without a starting cursor.
func (k *r2z2) run(ctx context.Context) {
	var seq int64
	for {
		if ctx.Err() != nil {
			return
		}
		s, err := k.fetchSequence(ctx)
		if err != nil {
			k.log.Errorw("failed to bootstrap r2z2 sequence", "error", err)
			if !sleepCtx(ctx, k.idle) {
				return
			}
			continue
		}
		seq = s
		break
	}

	k.log.Infow("r2z2 tail starting", "sequence", seq)

	for {
		if ctx.Err() != nil {
			return
		}

		km, status, err := k.fetchKillmail(ctx, seq)
		if err != nil {
			k.log.Errorw("failed to fetch r2z2 killmail", "sequence", seq, "error", err)
			if !sleepCtx(ctx, k.idle) {
				return
			}
			continue
		}

		switch status {
		case http.StatusOK:
			if km.ESI == nil || km.Zkb == nil {
				k.log.Warnw("malformed r2z2 killmail", "sequence", seq, "killmail_id", km.KillmailID)
				seq++
				continue
			}
			if k.seen.seen(km.KillmailID) {
				seq++
				continue
			}
			k.seen.add(km.KillmailID)
			select {
			case k.out <- km:
			case <-ctx.Done():
				return
			}
			seq++
			if !sleepCtx(ctx, k.active) {
				return
			}
		case http.StatusNotFound:
			if !sleepCtx(ctx, k.idle) {
				return
			}
		case http.StatusTooManyRequests:
			k.log.Warnw("r2z2 rate limited", "sequence", seq)
			if !sleepCtx(ctx, k.idle) {
				return
			}
		case http.StatusForbidden:
			k.log.Errorw("r2z2 forbidden (blank user-agent or ban)", "sequence", seq)
			if !sleepCtx(ctx, k.idle) {
				return
			}
		default:
			k.log.Errorw("unexpected r2z2 status", "sequence", seq, "status", status)
			if !sleepCtx(ctx, k.idle) {
				return
			}
		}
	}
}

type sequenceResponse struct {
	Sequence int64 `json:"sequence"`
}

// fetchSequence reports the latest known sequence id from
// /ephemeral/sequence.json.
func (k *r2z2) fetchSequence(ctx context.Context) (int64, error) {
	url := k.baseURL + "/ephemeral/sequence.json"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, fmt.Errorf("build sequence request: %w", err)
	}
	req.Header.Set("User-Agent", k.userAgent)

	res, err := k.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("fetch sequence: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("fetch sequence: got status %d", res.StatusCode)
	}

	var sr sequenceResponse
	if err := json.NewDecoder(res.Body).Decode(&sr); err != nil {
		return 0, fmt.Errorf("decode sequence: %w", err)
	}
	return sr.Sequence, nil
}

// fetchKillmail retrieves the killmail file at seq. On a non-200 response it
// returns (nil, status, nil) so the caller can branch on the status code; a
// transport or decode failure returns a non-nil error.
func (k *r2z2) fetchKillmail(ctx context.Context, seq int64) (*Killmail, int, error) {
	url := k.baseURL + "/ephemeral/" + strconv.FormatInt(seq, 10) + ".json"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, 0, fmt.Errorf("build killmail request: %w", err)
	}
	req.Header.Set("User-Agent", k.userAgent)

	res, err := k.client.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("fetch killmail: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, nil
	}

	km := &Killmail{}
	if err := json.NewDecoder(res.Body).Decode(km); err != nil {
		return nil, res.StatusCode, fmt.Errorf("decode killmail: %w", err)
	}
	return km, res.StatusCode, nil
}

// sleepCtx blocks for d or until ctx is cancelled. It reports false if ctx was
// cancelled, signalling the caller to stop.
func sleepCtx(ctx context.Context, d time.Duration) bool {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-t.C:
		return true
	case <-ctx.Done():
		return false
	}
}
