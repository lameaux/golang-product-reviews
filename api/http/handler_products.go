package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lameaux/golang-product-reviews/dto"
)

func (s *Server) setupProductsRouter(r *mux.Router) {
	r.HandleFunc("", s.handleListProducts()).Methods("GET")
	r.HandleFunc("/{product_id}", s.handleGetProduct()).Methods("GET")
	r.HandleFunc("", s.handlePostProduct()).Methods("POST")
	r.HandleFunc("/{product_id}", s.handlePutProduct()).Methods("PUT")
	r.HandleFunc("/{product_id}", s.handleDeleteProduct()).Methods("DELETE")
}

func (s *Server) handleListProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		offset, err := getIntQuery(r, "offset", 0)
		if err != nil {
			http.Error(w, "handleListProducts - invalid offset", http.StatusBadRequest)
			return
		}

		limit, err := getIntQuery(r, "limit", 100)
		if err != nil {
			http.Error(w, "handleListProducts - invalid limit", http.StatusBadRequest)
			return
		}

		products, err := s.manager.ListProducts(r.Context(), offset, limit)
		if err != nil {
			http.Error(w, "handleListProducts - ListProducts: "+err.Error(), http.StatusInternalServerError)
			return
		}

		s.sendAsJSON(w, products)
	}
}

func (s *Server) handleGetProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productID, err := getProductID(r)
		if err != nil {
			http.Error(w, "handleGetProduct - getProductID: "+err.Error(), http.StatusBadRequest)
			return
		}

		product, err := s.manager.GetProduct(r.Context(), productID)
		if err != nil {
			http.Error(w, "handleGetProduct - GetProduct: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if product == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		s.sendAsJSON(w, product)
	}
}

func (s *Server) handlePostProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var product dto.Product
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&product); err != nil {
			http.Error(w, "handlePostProduct - decode: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err := validate.Struct(&product); err != nil {
			http.Error(w, "handlePostProduct - validate: "+err.Error(), http.StatusBadRequest)
			return
		}

		productID, err := s.manager.CreateProduct(r.Context(), &product)
		if err != nil {
			http.Error(w, "handlePostProduct - CreateProduct: "+err.Error(), http.StatusInternalServerError)
			return
		}

		location := fmt.Sprintf("/products/%d", productID)
		w.Header().Add("Location", location)
		w.WriteHeader(http.StatusCreated)
	}
}

func (s *Server) handlePutProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var product dto.Product
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&product); err != nil {
			http.Error(w, "handlePutProduct - decode: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err := validate.Struct(&product); err != nil {
			http.Error(w, "handlePutProduct - validate: "+err.Error(), http.StatusBadRequest)
			return
		}

		productID, err := getProductID(r)
		if err != nil {
			http.Error(w, "handlePutProduct - getProductID: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err := s.manager.UpdateProduct(r.Context(), productID, &product); err != nil {
			http.Error(w, "handlePutProduct - manager: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) handleDeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productID, err := getProductID(r)
		if err != nil {
			http.Error(w, "handleDeleteProduct - getProductID: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err := s.manager.DeleteProduct(r.Context(), productID); err != nil {
			http.Error(w, "handleDeleteProduct - manager: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
