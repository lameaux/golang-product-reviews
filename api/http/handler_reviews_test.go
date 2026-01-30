package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandleListReviews(t *testing.T) {
	tests := []struct {
		name       string
		offset     string
		limit      string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "invalid offset",
			offset:     "invalid",
			wantStatus: http.StatusBadRequest,
			wantBody:   "invalid offset",
		},
		{
			name:       "invalid limit",
			limit:      "invalid",
			wantStatus: http.StatusBadRequest,
			wantBody:   "invalid limit",
		},
		{
			name:       "valid response",
			wantStatus: http.StatusOK,
			wantBody:   `[{"id":1,"first_name":"Sergej","last_name":"Sizov","review":"Perfect","rating":5}]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/products/1/reviews?offset=%s&limit=%s", tt.offset, tt.limit), nil)
			rec := httptest.NewRecorder()
			testRouter().ServeHTTP(rec, req)

			require.Equal(t, tt.wantStatus, rec.Code)
			require.Equal(t, tt.wantBody, strings.TrimSpace(rec.Body.String()))
		})
	}
}

func TestHandleGetReview(t *testing.T) {
	tests := []struct {
		name       string
		productID  int
		reviewID   int
		wantStatus int
		wantBody   string
	}{
		{
			name:       "invalid productID",
			productID:  404,
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "invalid reviewID",
			reviewID:   404,
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "valid id",
			productID:  1,
			reviewID:   1,
			wantStatus: http.StatusOK,
			wantBody:   `{"id":1,"first_name":"Sergej","last_name":"Sizov","review":"Perfect","rating":5}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/products/%d/reviews/%d", tt.productID, tt.reviewID), nil)
			rec := httptest.NewRecorder()
			testRouter().ServeHTTP(rec, req)

			require.Equal(t, tt.wantStatus, rec.Code)
			require.Equal(t, tt.wantBody, strings.TrimSpace(rec.Body.String()))
		})
	}
}

func TestHandlePostReview(t *testing.T) {
	tests := []struct {
		name         string
		body         string
		wantStatus   int
		wantBody     string
		wantLocation string
	}{
		{
			name:       "missing body",
			wantStatus: http.StatusBadRequest,
			wantBody:   "handlePostReview - decode",
		},
		{
			name:       "empty body",
			body:       "{}",
			wantStatus: http.StatusBadRequest,
			wantBody:   "handlePostReview - validate",
		},
		{
			name:         "valid",
			body:         `{"first_name":"John","last_name":"Doe","review":"Meh","rating":1}`,
			wantStatus:   http.StatusCreated,
			wantLocation: "/products/1/reviews/2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/products/1/reviews", strings.NewReader(tt.body))
			rec := httptest.NewRecorder()
			testRouter().ServeHTTP(rec, req)

			require.Equal(t, tt.wantStatus, rec.Code)
			require.Contains(t, strings.TrimSpace(rec.Body.String()), tt.wantBody)
			require.Equal(t, tt.wantLocation, rec.Header().Get("Location"))
		})
	}
}

func TestHandlePutReview(t *testing.T) {
	tests := []struct {
		name       string
		productID  int
		reviewID   int
		body       string
		wantStatus int
	}{
		{
			name:       "missing body",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "empty body",
			body:       "{}",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "invalid product id",
			productID:  404,
			body:       `{"first_name":"John","last_name":"Doe","review":"Meh","rating":1}`,
			wantStatus: http.StatusInternalServerError,
		},
		{
			name:       "invalid review id",
			reviewID:   404,
			body:       `{"first_name":"John","last_name":"Doe","review":"Meh","rating":1}`,
			wantStatus: http.StatusInternalServerError,
		},
		{
			name:       "valid",
			productID:  1,
			reviewID:   1,
			body:       `{"first_name":"John","last_name":"Doe","review":"Meh","rating":1}`,
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/products/%d/reviews/%d", tt.productID, tt.reviewID), strings.NewReader(tt.body))
			rec := httptest.NewRecorder()
			testRouter().ServeHTTP(rec, req)

			require.Equal(t, tt.wantStatus, rec.Code)
		})
	}
}

func TestHandleDeleteReview(t *testing.T) {
	tests := []struct {
		name       string
		productID  int
		reviewID   int
		wantStatus int
	}{
		{
			name:       "invalid product id",
			productID:  404,
			wantStatus: http.StatusInternalServerError,
		},
		{
			name:       "invalid review id",
			reviewID:   404,
			wantStatus: http.StatusInternalServerError,
		},
		{
			name:       "valid",
			productID:  1,
			reviewID:   1,
			wantStatus: http.StatusNoContent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/products/%d/reviews/%d", tt.productID, tt.reviewID), nil)
			rec := httptest.NewRecorder()
			testRouter().ServeHTTP(rec, req)

			require.Equal(t, tt.wantStatus, rec.Code)
		})
	}
}
