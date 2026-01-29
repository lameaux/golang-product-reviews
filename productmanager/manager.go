package productmanager

import (
	"github.com/lameaux/golang-product-reviews/database"
	"github.com/lameaux/golang-product-reviews/dto"
	"github.com/lameaux/golang-product-reviews/model"
)

type Manager struct {
	dao database.DAO
}

func New(dao database.DAO) *Manager {
	return &Manager{dao: dao}
}

func (*Manager) CreateProduct(p *dto.Product) (model.ID, error) {
	return 0, nil
}

func (*Manager) UpdateProduct(productID model.ID, p *dto.Product) error {
	return nil
}

func (*Manager) DeleteProduct(productID model.ID) error {
	return nil
}

func (*Manager) GetProduct(productID model.ID) *dto.ProductWithRating {
	return nil
}

func (*Manager) ListProducts(offset int, limit int) []*dto.ProductWithRating {
	return []*dto.ProductWithRating{}
}

func (*Manager) CreateProductReview(productID model.ID, r *dto.Review) (model.ID, error) {
	return 0, nil
}

func (*Manager) DeleteProductReview(productID model.ID, reviewID model.ID) error {
	return nil
}

func (*Manager) UpdateProductReview(productID model.ID, reviewID model.ID, review *dto.Review) error {
	return nil
}

func (*Manager) GetProductReview(productID model.ID, reviewID model.ID) *dto.Review {
	return nil
}

func (*Manager) ListProductReviews(productID model.ID, offset int, limit int) []*dto.Review {
	return []*dto.Review{}
}
