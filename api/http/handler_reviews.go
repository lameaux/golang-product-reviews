package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lameaux/golang-product-reviews/dto"
)

func (s *Server) setupReviewsRouter(r *mux.Router) {
	r.HandleFunc("", s.handleListReviews()).Methods("GET")
	r.HandleFunc("/{review_id}", s.handleGetReview()).Methods("GET")
	r.HandleFunc("", s.handlePostReview()).Methods("POST")
	r.HandleFunc("/{review_id}", s.handlePutReview()).Methods("PUT")
	r.HandleFunc("/{review_id}", s.handleDeleteReview()).Methods("DELETE")
}

func (s *Server) handleListReviews() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		offset, err := getIntQuery(r, "offset", 0)
		if err != nil {
			http.Error(w, "invalid offset", http.StatusBadRequest)
			return
		}

		limit, err := getIntQuery(r, "limit", 100)
		if err != nil {
			http.Error(w, "invalid limit", http.StatusBadRequest)
			return
		}

		productID, err := getProductID(r)
		if err != nil {
			http.Error(w, "handleListReviews - getProductID: "+err.Error(), http.StatusBadRequest)
			return
		}

		reviews := s.manager.ListProductReviews(productID, offset, limit)
		s.sendAsJSON(w, reviews)
	}
}

func (s *Server) handleGetReview() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productID, err := getProductID(r)
		if err != nil {
			http.Error(w, "handleGetReview - getProductID: "+err.Error(), http.StatusBadRequest)
			return
		}

		reviewID, err := getReviewID(r)
		if err != nil {
			http.Error(w, "handleGetReview - getReviewID: "+err.Error(), http.StatusBadRequest)
			return
		}

		review := s.manager.GetProductReview(productID, reviewID)
		if review == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		s.sendAsJSON(w, review)
	}
}

func (s *Server) handlePostReview() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		productID, err := getProductID(r)
		if err != nil {
			http.Error(w, "handlePostReview - getProductID: "+err.Error(), http.StatusBadRequest)
			return
		}

		var review dto.Review
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&review); err != nil {
			http.Error(w, "handlePostReview - decode: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err := validate.Struct(&review); err != nil {
			http.Error(w, "handlePostReview - validate: "+err.Error(), http.StatusBadRequest)
			return
		}

		reviewID, err := s.manager.CreateProductReview(productID, &review)
		if err != nil {
			http.Error(w, "handlePostReview - CreateProductReview: "+err.Error(), http.StatusInternalServerError)
			return
		}

		location := fmt.Sprintf("/products/%d/reviews/%d", productID, reviewID)
		w.Header().Add("Location", location)
		w.WriteHeader(http.StatusCreated)
	}
}

func (s *Server) handlePutReview() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		productID, err := getProductID(r)
		if err != nil {
			http.Error(w, "handlePutReview - getProductID: "+err.Error(), http.StatusBadRequest)
			return
		}

		reviewID, err := getReviewID(r)
		if err != nil {
			http.Error(w, "handlePutReview - getReviewID: "+err.Error(), http.StatusBadRequest)
			return
		}

		var review dto.Review
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&review); err != nil {
			http.Error(w, "handlePutReview - decode: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err := validate.Struct(&review); err != nil {
			http.Error(w, "handlePutReview - validate: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err := s.manager.UpdateProductReview(productID, reviewID, &review); err != nil {
			http.Error(w, "handlePutReview - manager: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) handleDeleteReview() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productID, err := getProductID(r)
		if err != nil {
			http.Error(w, "handleDeleteReview - getProductID: "+err.Error(), http.StatusBadRequest)
			return
		}

		reviewID, err := getReviewID(r)
		if err != nil {
			http.Error(w, "handleDeleteReview - getReviewID: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err := s.manager.DeleteProductReview(productID, reviewID); err != nil {
			http.Error(w, "handleDeleteReview - manager: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
