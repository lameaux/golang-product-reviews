package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/lameaux/golang-product-reviews/productmanager"
	"github.com/rs/zerolog"
)

type Server struct {
	srv     *http.Server
	port    int
	logger  zerolog.Logger
	manager *productmanager.Manager
}

func New(
	port int,
	logger zerolog.Logger,
	manager *productmanager.Manager,
) *Server {
	return &Server{
		port:    port,
		logger:  logger,
		manager: manager,
		srv: &http.Server{
			Addr:         fmt.Sprintf(":%d", port),
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
		},
	}
}

func (s *Server) Serve() error {
	s.srv.Handler = s.CreateRouter()

	s.logger.Info().Int("port", s.port).Msg("starting http server")
	err := s.srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		s.logger.Error().Int("port", s.port).Err(err).Msg("http server failed to start")
		return err
	}

	return nil
}

func (s *Server) CreateRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", s.handleHealth()).Methods("GET")

	return r
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		s.logger.Error().Err(err).Msg("HTTP server shutdown error")
	}
}
