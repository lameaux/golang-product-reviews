package productmanager

import (
	"context"
	"errors"

	"github.com/lameaux/golang-product-reviews/dto"
	"github.com/lameaux/golang-product-reviews/model"
)

var _ Manager = (*StubManager)(nil)

type StubManager struct {
	Products []*dto.ProductWithRating
	Reviews  []*dto.Review
}

func (s *StubManager) CreateProduct(ctx context.Context, p *dto.Product) (model.ID, error) {
	return len(s.Products) + 1, nil
}

func (s *StubManager) UpdateProduct(ctx context.Context, productID model.ID, p *dto.Product) error {
	if productID > len(s.Products) {
		return errors.New("not found")
	}

	return nil
}

func (s *StubManager) DeleteProduct(ctx context.Context, productID model.ID) error {
	if productID > len(s.Products) {
		return errors.New("not found")
	}

	return nil
}

func (s *StubManager) GetProduct(ctx context.Context, productID model.ID) (*dto.ProductWithRating, error) {
	if productID > len(s.Products) {
		return nil, nil
	}

	return s.Products[productID-1], nil
}

func (s *StubManager) ListProducts(context.Context, int, int) ([]*dto.ProductWithRating, error) {
	return s.Products, nil
}

func (s *StubManager) CreateProductReview(ctx context.Context, productID model.ID, r *dto.Review) (model.ID, error) {
	return len(s.Reviews) + 1, nil
}

func (s *StubManager) DeleteProductReview(ctx context.Context, productID model.ID, reviewID model.ID) error {
	if productID > len(s.Products) {
		return errors.New("not found")
	}

	if reviewID > len(s.Reviews) {
		return errors.New("not found")
	}

	return nil
}

func (s *StubManager) UpdateProductReview(ctx context.Context, productID model.ID, reviewID model.ID, review *dto.Review) error {
	if productID > len(s.Products) {
		return errors.New("not found")
	}

	if reviewID > len(s.Reviews) {
		return errors.New("not found")
	}

	return nil
}

func (s *StubManager) GetProductReview(ctx context.Context, productID model.ID, reviewID model.ID) (*dto.Review, error) {
	if productID > len(s.Products) {
		return nil, nil
	}

	if reviewID > len(s.Reviews) {
		return nil, nil
	}

	return s.Reviews[reviewID-1], nil
}

func (s *StubManager) ListProductReviews(ctx context.Context, productID model.ID, offset int, limit int) ([]*dto.Review, error) {
	return s.Reviews, nil
}
