package observer

import (
	"context"
	"net/http"
	"sync"

	"github.com/antihax/goesi"
	"go.uber.org/zap"
)

type Observer struct {
	poller *r2z2
	router *router
	sender *sender
	wg     sync.WaitGroup
}

// New wires the poller, router, and sender without starting any goroutines.
// Call Start to begin processing and Wait to block until shutdown completes.
func New(config ObserverConfig, log *zap.SugaredLogger) (*Observer, error) {
	parsedDests, err := GetDestinations(config)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{Timeout: httpTimeout}
	esiClient := goesi.NewAPIClient(httpClient, config.EsiUserAgent)

	killmailChan := make(chan *Killmail)
	routerChan := make(chan *RoutedKillmail)

	return &Observer{
		poller: newR2Z2(killmailChan, httpClient, r2z2BaseURL, config.EsiUserAgent, config.ActivePollInterval, config.IdlePollInterval, log),
		router: newRouter(killmailChan, routerChan, parsedDests, log),
		sender: newSender(routerChan, esiClient, log),
	}, nil
}

// Start launches the poller, router, and sender goroutines. They run until ctx
// is cancelled.
func (o *Observer) Start(ctx context.Context) {
	o.wg.Add(3)
	go func() {
		defer o.wg.Done()
		o.poller.run(ctx)
	}()
	go func() {
		defer o.wg.Done()
		o.router.run(ctx)
	}()
	go func() {
		defer o.wg.Done()
		o.sender.run(ctx)
	}()
}

// Wait blocks until all goroutines started by Start have returned.
func (o *Observer) Wait() {
	o.wg.Wait()
}
