package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	httpapi "github.com/lameaux/golang-product-reviews/api/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	logger.Info().Msg("api started")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	httpPort, err := strconv.Atoi(port)
	if err != nil {
		logger.Error().Err(err).Msg("invalid PORT")
		return
	}

	httpServer := httpapi.New(httpPort, logger)
	httpErrCh := make(chan error, 1)
	go func() {
		httpErrCh <- httpServer.Serve()
	}()
	defer httpServer.Stop()

	<-ctx.Done()
	logger.Info().Msg("api shutting down")
}
