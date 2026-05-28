package observer

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"go.uber.org/zap"
)

const testUserAgent = "zobserver-test/maintainer@example.com"

// killmailFixture renders a minimal but structurally complete R2Z2 killmail
// file: one victim, one attacker with the final blow, and zkb metadata.
func killmailFixture(killmailID, sequenceID int64, totalValue float64) string {
	return fmt.Sprintf(`{
		"killmail_id": %d,
		"hash": "abc123",
		"esi": {
			"killmail_id": %d,
			"killmail_time": "2026-05-28T22:08:55Z",
			"solar_system_id": 30001388,
			"victim": {
				"alliance_id": 99000001,
				"character_id": 90000001,
				"corporation_id": 98000001,
				"damage_taken": 5000,
				"ship_type_id": 28710
			},
			"attackers": [
				{
					"alliance_id": 99000002,
					"character_id": 90000002,
					"corporation_id": 98000002,
					"damage_done": 5000,
					"final_blow": true,
					"ship_type_id": 17738
				}
			]
		},
		"zkb": {
			"locationID": 50001074,
			"hash": "abc123",
			"totalValue": %f,
			"points": 7,
			"npc": false
		},
		"uploaded_at": 1780006305,
		"sequence_id": %d
	}`, killmailID, killmailID, totalValue, sequenceID)
}

func testLogger(t *testing.T) *zap.SugaredLogger {
	t.Helper()
	return zap.NewNop().Sugar()
}

// recordingMux records every inbound User-Agent so tests can assert the header
// is always set to the configured value.
type recordingMux struct {
	mu         sync.Mutex
	userAgents []string
}

func (m *recordingMux) record(r *http.Request) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.userAgents = append(m.userAgents, r.Header.Get("User-Agent"))
}

func (m *recordingMux) seenUserAgents() []string {
	m.mu.Lock()
	defer m.mu.Unlock()
	out := make([]string, len(m.userAgents))
	copy(out, m.userAgents)
	return out
}

func (m *recordingMux) assertUserAgents(t *testing.T) {
	t.Helper()
	seen := m.seenUserAgents()
	if len(seen) == 0 {
		t.Errorf("expected at least one recorded request")
	}
	for _, ua := range seen {
		if ua == "" {
			t.Errorf("user-agent: got blank, want %q", testUserAgent)
		}
		if ua != testUserAgent {
			t.Errorf("user-agent: got %q, want %q", ua, testUserAgent)
		}
	}
}

func newTestPoller(t *testing.T, server *httptest.Server, out chan *Killmail) *r2z2 {
	t.Helper()
	return newR2Z2(out, server.Client(), server.URL, testUserAgent, 1*time.Millisecond, 5*time.Millisecond, testLogger(t))
}

func TestFetchSequence(t *testing.T) {
	rec := &recordingMux{}
	mux := http.NewServeMux()
	mux.HandleFunc("/ephemeral/sequence.json", func(w http.ResponseWriter, r *http.Request) {
		rec.record(r)
		_, _ = w.Write([]byte(`{"sequence":97719373}`))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	poller := newTestPoller(t, server, make(chan *Killmail))
	seq, err := poller.fetchSequence(context.Background())
	if err != nil {
		t.Fatalf("fetchSequence: unexpected error: %v", err)
	}
	if seq != 97719373 {
		t.Errorf("sequence: got %d, want %d", seq, 97719373)
	}

	rec.assertUserAgents(t)
}

func TestFetchKillmail(t *testing.T) {
	testCases := []struct {
		name           string
		status         int
		body           string
		wantStatus     int
		wantKillmail   bool
		wantKillmailID int32
		wantTotalValue float64
	}{
		{
			name:           "OK",
			status:         http.StatusOK,
			body:           killmailFixture(135834513, 97719373, 1546276296.29),
			wantStatus:     http.StatusOK,
			wantKillmail:   true,
			wantKillmailID: 135834513,
			wantTotalValue: 1546276296.29,
		},
		{
			name:       "NotFound",
			status:     http.StatusNotFound,
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "Forbidden",
			status:     http.StatusForbidden,
			wantStatus: http.StatusForbidden,
		},
		{
			name:       "TooManyRequests",
			status:     http.StatusTooManyRequests,
			wantStatus: http.StatusTooManyRequests,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := &recordingMux{}
			mux := http.NewServeMux()
			mux.HandleFunc("/ephemeral/", func(w http.ResponseWriter, r *http.Request) {
				rec.record(r)
				if tc.status != http.StatusOK {
					w.WriteHeader(tc.status)
					return
				}
				_, _ = w.Write([]byte(tc.body))
			})
			server := httptest.NewServer(mux)
			defer server.Close()

			poller := newTestPoller(t, server, make(chan *Killmail))
			km, status, err := poller.fetchKillmail(context.Background(), 97719373)
			if err != nil {
				t.Fatalf("fetchKillmail: unexpected error: %v", err)
			}
			if status != tc.wantStatus {
				t.Errorf("status: got %d, want %d", status, tc.wantStatus)
			}

			if !tc.wantKillmail {
				if km != nil {
					t.Errorf("killmail: got non-nil, want nil")
				}
				rec.assertUserAgents(t)
				return
			}

			if km == nil {
				t.Fatalf("killmail: got nil, want non-nil")
			}
			if km.ESI == nil {
				t.Fatalf("killmail ESI: got nil, want non-nil")
			}
			if km.ESI.KillmailId != tc.wantKillmailID {
				t.Errorf("ESI.KillmailId: got %d, want %d", km.ESI.KillmailId, tc.wantKillmailID)
			}
			if km.ESI.Victim.ShipTypeId != 28710 {
				t.Errorf("ESI.Victim.ShipTypeId: got %d, want %d", km.ESI.Victim.ShipTypeId, 28710)
			}
			if len(km.ESI.Attackers) != 1 {
				t.Fatalf("ESI.Attackers: got %d, want 1", len(km.ESI.Attackers))
			}
			if km.Zkb == nil {
				t.Fatalf("killmail Zkb: got nil, want non-nil")
			}
			if km.Zkb.TotalValue != tc.wantTotalValue {
				t.Errorf("Zkb.TotalValue: got %f, want %f", km.Zkb.TotalValue, tc.wantTotalValue)
			}

			rec.assertUserAgents(t)
		})
	}
}

func TestRun(t *testing.T) {
	const bootSeq int64 = 100
	rec := &recordingMux{}

	mux := http.NewServeMux()
	mux.HandleFunc("/ephemeral/sequence.json", func(w http.ResponseWriter, r *http.Request) {
		rec.record(r)
		_, _ = fmt.Fprintf(w, `{"sequence":%d}`, bootSeq)
	})
	mux.HandleFunc("/ephemeral/", func(w http.ResponseWriter, r *http.Request) {
		rec.record(r)
		switch sequenceFromPath(t, r.URL.Path) {
		case 100:
			_, _ = w.Write([]byte(killmailFixture(111, 100, 10.0)))
		case 101:
			// Same killmail_id under a new sequence id: dedup must drop it.
			_, _ = w.Write([]byte(killmailFixture(111, 101, 10.0)))
		case 102:
			_, _ = w.Write([]byte(killmailFixture(222, 102, 20.0)))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	out := make(chan *Killmail)
	poller := newTestPoller(t, server, out)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var mu sync.Mutex
	var collected []*Killmail
	var collectWG sync.WaitGroup
	collectWG.Go(func() {
		for km := range out {
			mu.Lock()
			collected = append(collected, km)
			mu.Unlock()
		}
	})

	done := make(chan struct{})
	go func() {
		poller.run(ctx)
		close(done)
	}()

	collectedLen := func() int {
		mu.Lock()
		defer mu.Unlock()
		return len(collected)
	}

	deadline := time.After(3 * time.Second)
	for collectedLen() < 2 {
		select {
		case <-deadline:
			t.Fatalf("timed out waiting for emitted killmails, got %d", collectedLen())
		case <-time.After(2 * time.Millisecond):
		}
	}

	cancel()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatalf("run did not return promptly after cancel")
	}
	close(out)
	collectWG.Wait()

	if len(collected) != 2 {
		t.Fatalf("emitted killmails: got %d, want 2 (dedup of recurring id failed)", len(collected))
	}
	if collected[0].KillmailID != 111 {
		t.Errorf("first killmail id: got %d, want 111", collected[0].KillmailID)
	}
	if collected[1].KillmailID != 222 {
		t.Errorf("second killmail id: got %d, want 222", collected[1].KillmailID)
	}
	if collected[1].Zkb.TotalValue != 20.0 {
		t.Errorf("second killmail total value: got %f, want 20.0", collected[1].Zkb.TotalValue)
	}

	rec.assertUserAgents(t)
}

func TestSeenSet(t *testing.T) {
	s := newSeenSet(2)
	if s.seen(1) {
		t.Errorf("seen(1): got true on empty set, want false")
	}
	s.add(1)
	s.add(2)
	if !s.seen(1) || !s.seen(2) {
		t.Errorf("expected 1 and 2 to be present")
	}

	s.add(3)
	if s.seen(1) {
		t.Errorf("seen(1): got true after eviction, want false")
	}
	if !s.seen(2) || !s.seen(3) {
		t.Errorf("expected 2 and 3 to be present after eviction")
	}

	s.add(2)
	if len(s.order) != 2 {
		t.Errorf("order length after re-adding existing id: got %d, want 2", len(s.order))
	}
}

func sequenceFromPath(t *testing.T, path string) int64 {
	t.Helper()
	base := strings.TrimSuffix(strings.TrimPrefix(path, "/ephemeral/"), ".json")
	seq, err := strconv.ParseInt(base, 10, 64)
	if err != nil {
		t.Fatalf("parse sequence from path %q: %v", path, err)
	}
	return seq
}
