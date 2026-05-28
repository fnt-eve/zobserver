package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/fnt-eve/zobserver/internal/logger"
	"github.com/fnt-eve/zobserver/internal/observer"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	logger.LoggerConfig
	observer.ObserverConfig
}

func main() {
	// Parse env vars
	var c config
	err := envconfig.Process("", &c)
	if err != nil {
		panic(err)
	}
	// Init logger
	logger, err := c.InitLogger()
	if err != nil {
		panic(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	obs, err := observer.New(c.ObserverConfig, logger.Sugar())
	if err != nil {
		logger.Sugar().Errorw("failed to start km feed", "error", err)
		os.Exit(1)
	}

	obs.Start(ctx)
	logger.Sugar().Infow("observer running")

	<-ctx.Done()
	stop()
	obs.Wait()
}
