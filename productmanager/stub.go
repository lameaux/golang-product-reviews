package productmanager

import (
	"errors"

	"github.com/lameaux/golang-product-reviews/dto"
	"github.com/lameaux/golang-product-reviews/model"
)

var _ Manager = (*StubManager)(nil)

type StubManager struct {
	Products []*dto.ProductWithRating
	Reviews  []*dto.Review
}

func (s *StubManager) CreateProduct(p *dto.Product) (model.ID, error) {
	return len(s.Products) + 1, nil
}

func (s *StubManager) UpdateProduct(productID model.ID, p *dto.Product) error {
	if productID > len(s.Products) {
		return errors.New("not found")
	}

	return nil
}

func (s *StubManager) DeleteProduct(productID model.ID) error {
	if productID > len(s.Products) {
		return errors.New("not found")
	}

	return nil
}

func (s *StubManager) GetProduct(productID model.ID) *dto.ProductWithRating {
	if productID > len(s.Products) {
		return nil
	}

	return s.Products[productID-1]
}

func (s *StubManager) ListProducts(int, int) []*dto.ProductWithRating {
	return s.Products
}

func (*StubManager) CreateProductReview(productID model.ID, r *dto.Review) (model.ID, error) {
	return 0, nil
}

func (*StubManager) DeleteProductReview(productID model.ID, reviewID model.ID) error {
	return nil
}

func (*StubManager) UpdateProductReview(productID model.ID, reviewID model.ID, review *dto.Review) error {
	return nil
}

func (*StubManager) GetProductReview(productID model.ID, reviewID model.ID) *dto.Review {
	return nil
}

func (*StubManager) ListProductReviews(productID model.ID, offset int, limit int) []*dto.Review {
	return []*dto.Review{}
}
