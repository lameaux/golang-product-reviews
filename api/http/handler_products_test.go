package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandleListProducts(t *testing.T) {
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
			wantBody:   "handleListProducts - invalid offset",
		},
		{
			name:       "invalid limit",
			limit:      "invalid",
			wantStatus: http.StatusBadRequest,
			wantBody:   "handleListProducts - invalid limit",
		},
		{
			name:       "valid response",
			wantStatus: http.StatusOK,
			wantBody:   `[{"id":1,"name":"P1","description":"P1 desc","price":100,"rating":1}]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/products?offset=%s&limit=%s", tt.offset, tt.limit), nil)
			rec := httptest.NewRecorder()
			testRouter().ServeHTTP(rec, req)

			require.Equal(t, tt.wantStatus, rec.Code)
			require.Equal(t, tt.wantBody, strings.TrimSpace(rec.Body.String()))
		})
	}
}

func TestHandleGetProduct(t *testing.T) {
	tests := []struct {
		name       string
		id         int
		wantStatus int
		wantBody   string
	}{
		{
			name:       "invalid id",
			id:         404,
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "valid id",
			id:         1,
			wantStatus: http.StatusOK,
			wantBody:   `{"id":1,"name":"P1","description":"P1 desc","price":100,"rating":1}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/products/%d", tt.id), nil)
			rec := httptest.NewRecorder()
			testRouter().ServeHTTP(rec, req)

			require.Equal(t, tt.wantStatus, rec.Code)
			require.Equal(t, tt.wantBody, strings.TrimSpace(rec.Body.String()))
		})
	}
}

func TestHandlePostProduct(t *testing.T) {
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
			wantBody:   "handlePostProduct - decode",
		},
		{
			name:       "empty body",
			body:       "{}",
			wantStatus: http.StatusBadRequest,
			wantBody:   "handlePostProduct - validate",
		},
		{
			name:         "valid",
			body:         `{"name":"P2","description":"P2 desc","price":200}`,
			wantStatus:   http.StatusCreated,
			wantLocation: "/products/2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(tt.body))
			rec := httptest.NewRecorder()
			testRouter().ServeHTTP(rec, req)

			require.Equal(t, tt.wantStatus, rec.Code)
			require.Contains(t, strings.TrimSpace(rec.Body.String()), tt.wantBody)
			require.Equal(t, tt.wantLocation, rec.Header().Get("Location"))
		})
	}
}

func TestHandlePutProduct(t *testing.T) {
	tests := []struct {
		name       string
		id         int
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
			name:       "invalid id",
			id:         404,
			body:       `{"name":"P2","description":"P2 desc","price":200}`,
			wantStatus: http.StatusInternalServerError,
		},
		{
			name:       "valid",
			id:         1,
			body:       `{"name":"P2","description":"P2 desc","price":200}`,
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/products/%d", tt.id), strings.NewReader(tt.body))
			rec := httptest.NewRecorder()
			testRouter().ServeHTTP(rec, req)

			require.Equal(t, tt.wantStatus, rec.Code)
		})
	}
}

func TestHandleDeleteProduct(t *testing.T) {
	tests := []struct {
		name       string
		id         int
		wantStatus int
	}{
		{
			name:       "invalid id",
			id:         404,
			wantStatus: http.StatusInternalServerError,
		},
		{
			name:       "valid",
			id:         1,
			wantStatus: http.StatusNoContent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/products/%d", tt.id), nil)
			rec := httptest.NewRecorder()
			testRouter().ServeHTTP(rec, req)

			require.Equal(t, tt.wantStatus, rec.Code)
		})
	}
}
