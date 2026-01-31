package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/nats-io/nats.go"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("audit starting")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	natsURL := os.Getenv("NATS_URL")

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Error().Str("url", natsURL).Err(err).Msg("NATS connect failed")
	}
	defer nc.Close()

	log.Info().Str("url", natsURL).Msg("connected to NATS")

	sub, err := nc.Subscribe("reviews", func(msg *nats.Msg) {
		log.Info().Str("data", string(msg.Data)).Msg("review updated")
	})
	if err != nil {
		log.Error().Err(err).Msg("subscribe failed")
	}
	defer sub.Unsubscribe()

	log.Info().Msg("audit started")

	<-ctx.Done()
	log.Info().Msg("audit stopped")
}
