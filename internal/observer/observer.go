package observer

import (
	"net/http"

	"github.com/antihax/goesi"
	"go.uber.org/zap"
)

type Observer struct {
	redisq *redisQ
	router *router
	sender *sender
}

func New(config ObserverConfig, log *zap.SugaredLogger) (*Observer, error) {
	parsedDests, err := GetDestinations(config)
	if err != nil {
		return nil, err
	}

	esiClient := goesi.NewAPIClient(http.DefaultClient, "config.EsiUserAgent")
	redisChan := make(chan *ZkilResponse)
	routerChan := make(chan *RoutedZkilResponse)
	return &Observer{
		redisq: newRedisQ(redisChan, config.QueueName, config.TTW, log),
		router: newRouter(redisChan, routerChan, parsedDests, log),
		sender: newSender(routerChan, esiClient, log),
	}, err
}
