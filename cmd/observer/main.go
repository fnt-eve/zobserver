package main

import (
	"fmt"
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

	_, err = observer.New(c.ObserverConfig, logger.Sugar())
	if err != nil {
		logger.Sugar().Errorw("failed to start km feed", "error", err)
		os.Exit(1)
	}

	fmt.Println("Server is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
