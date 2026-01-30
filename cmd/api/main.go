package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	httpapi "github.com/lameaux/golang-product-reviews/api/http"
	"github.com/lameaux/golang-product-reviews/database"
	"github.com/lameaux/golang-product-reviews/notifier"
	"github.com/lameaux/golang-product-reviews/productmanager"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	logger.Info().Msg("api starting")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := run(ctx, &logger); err != nil {
		logger.Error().Err(err).Msg("start failed")
	}

	logger.Info().Msg("api stopped")
}

func run(ctx context.Context, logger *zerolog.Logger) error {
	dao, err := setupDatabase()
	if err != nil {
		return fmt.Errorf("setupDatabase: %w", err)
	}

	reviewNotifier := notifier.New(logger)
	manager := productmanager.New(dao, reviewNotifier.Notify)

	httpPort, err := getHttpPort()
	if err != nil {
		return fmt.Errorf("invalid port: %w", err)
	}

	httpServer := httpapi.New(httpPort, logger, manager)

	httpErrCh := make(chan error, 1)
	go func() {
		httpErrCh <- httpServer.Serve()
	}()
	defer httpServer.Stop()

	logger.Info().Msg("api started")

	select {
	case <-ctx.Done():
		logger.Info().Msg("api shutting down")
	case err := <-httpErrCh:
		if err != nil {
			return fmt.Errorf("http server error: %w", err)
		}
		logger.Info().Msg("http server stopped")
	}

	return nil
}

func setupDatabase() (database.DAO, error) {
	gormDB, sqlDB, err := database.Connect(os.Getenv("POSTGRES_URL"))
	if err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}

	if os.Getenv("RUN_MIGRATIONS") == "true" {
		if err := database.Migrate(sqlDB); err != nil {
			return nil, fmt.Errorf("migrations failed: %w", err)
		}
	}

	return database.NewPostgresDAO(gormDB), nil
}

func getHttpPort() (int, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	httpPort, err := strconv.Atoi(port)
	if err != nil {
		return 0, fmt.Errorf("invalid port: %w", err)
	}

	return httpPort, nil
}
