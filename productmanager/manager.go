package productmanager

import (
	"github.com/lameaux/golang-product-reviews/dto"
	"github.com/lameaux/golang-product-reviews/model"
)

type Manager interface {
	CreateProduct(p *dto.Product) (model.ID, error)
	UpdateProduct(productID model.ID, p *dto.Product) error
	DeleteProduct(productID model.ID) error

	GetProduct(productID model.ID) *dto.ProductWithRating
	ListProducts(offset int, limit int) []*dto.ProductWithRating

	CreateProductReview(productID model.ID, r *dto.Review) (model.ID, error)
	DeleteProductReview(productID model.ID, reviewID model.ID) error
	UpdateProductReview(productID model.ID, reviewID model.ID, review *dto.Review) error

	GetProductReview(productID model.ID, reviewID model.ID) *dto.Review
	ListProductReviews(productID model.ID, offset int, limit int) []*dto.Review
}
