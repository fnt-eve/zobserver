package observer

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

var RedisQURL = "https://zkillredisq.stream/listen.php"

var errTooManyRequests = errors.New("too many requests")

type RequestError struct {
	StatusCode int

	Err error
}

func (r *RequestError) Error() string {
	return r.Err.Error()
}

type redisQ struct {
	out     chan *ZkilResponse
	queueID string
	ttw     string
}

func newRedisQ(c chan *ZkilResponse, queueID string, ttw string, log *zap.SugaredLogger) *redisQ {
	if queueID == "" {
		r, _ := GenRand(10)
		queueID = *r
	}
	k := &redisQ{c, queueID, ttw}
	k.init(log)
	return k
}

func (k *redisQ) init(log *zap.SugaredLogger) {
	go func() {
		for {
			log.Debugw("polling redisq API", "queueID", k.queueID, "ttw", k.ttw)
			resp, err := queryRedisq(k.queueID, k.ttw)
			// Delay if 429 error
			if err == errTooManyRequests {
				time.Sleep(1 * time.Second)
				continue
			}

			if err != nil {
				log.Errorw("error while fetching from RedisQ API", "error", err)
				time.Sleep(1 * time.Second)
				continue
			}

			// Skip empty response
			if resp.Package.KillID == 0 {
				continue
			}

			log.Debugw("fetched KM from RedisQ", "killID", resp.Package.KillID)

			k.out <- resp
		}
	}()
}

func queryRedisq(queueID string, ttw string) (*ZkilResponse, error) {
	req, err := http.NewRequest("GET", RedisQURL, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("queueID", queueID)
	q.Add("ttw", ttw)
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusTooManyRequests {
		return nil, errTooManyRequests
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("got %v from RedisQ", res.StatusCode)
	}

	zkr := &ZkilResponse{}
	err = json.NewDecoder(res.Body).Decode(zkr)
	if err != nil {
		return nil, err
	}

	return zkr, nil
}

func GenRand(length int) (*string, error) {
	// Generate a random string
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	state := base64.URLEncoding.EncodeToString(b)
	return &state, nil
}
