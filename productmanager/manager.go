package productmanager

import (
	"context"

	"github.com/lameaux/golang-product-reviews/dto"
	"github.com/lameaux/golang-product-reviews/model"
)

type Manager interface {
	CreateProduct(ctx context.Context, p *dto.Product) (model.ID, error)
	UpdateProduct(ctx context.Context, productID model.ID, p *dto.Product) error
	DeleteProduct(ctx context.Context, productID model.ID) error

	GetProduct(ctx context.Context, productID model.ID) (*dto.ProductWithRating, error)
	ListProducts(ctx context.Context, offset int, limit int) ([]*dto.ProductWithRating, error)

	CreateProductReview(ctx context.Context, productID model.ID, r *dto.Review) (model.ID, error)
	DeleteProductReview(ctx context.Context, productID model.ID, reviewID model.ID) error
	UpdateProductReview(ctx context.Context, productID model.ID, reviewID model.ID, review *dto.Review) error

	GetProductReview(ctx context.Context, productID model.ID, reviewID model.ID) (*dto.Review, error)
	ListProductReviews(ctx context.Context, productID model.ID, offset int, limit int) ([]*dto.Review, error)
}
