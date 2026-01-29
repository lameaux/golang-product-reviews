package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/lameaux/golang-product-reviews/model"
	"github.com/lameaux/golang-product-reviews/productmanager"
	"github.com/rs/zerolog"
)

var (
	validate = validator.New(validator.WithRequiredStructEnabled())
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
	r.Use(s.loggingMiddleware)
	r.HandleFunc("/health", s.handleHealth()).Methods("GET")

	products := r.PathPrefix("/products").Subrouter()
	s.setupProductsRouter(products)

	reviews := products.PathPrefix("/{product_id}/reviews").Subrouter()
	s.setupReviewsRouter(reviews)

	return r
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		s.logger.Error().Err(err).Msg("HTTP server shutdown error")
	}
}

func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info().Str("method", r.Method).Msg(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func (s *Server) sendAsJSON(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		s.logger.Error().Err(err).Msg("encode response failed")
	}
}

func getProductID(r *http.Request) (model.ID, error) {
	productID, err := strconv.Atoi(mux.Vars(r)["product_id"])
	if err != nil {
		return 0, err
	}

	return productID, nil
}

func getReviewID(r *http.Request) (model.ID, error) {
	reviewID, err := strconv.Atoi(mux.Vars(r)["review_id"])
	if err != nil {
		return 0, err
	}

	return reviewID, nil
}

func getIntQuery(r *http.Request, key string, def int) (int, error) {
	val := r.URL.Query().Get(key)
	if val == "" {
		return def, nil
	}
	return strconv.Atoi(val)
}
