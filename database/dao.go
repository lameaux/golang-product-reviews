package database

import (
	"context"

	"github.com/lameaux/golang-product-reviews/model"
)

type DAO interface {
	CreateProduct(ctx context.Context, p *model.Product) (model.ID, error)
	UpdateProduct(ctx context.Context, p *model.Product) error
	DeleteProduct(ctx context.Context, id model.ID) error
	GetProduct(ctx context.Context, id model.ID) (*model.Product, error)
	GetProductRating(ctx context.Context, id model.ID) (float32, error)
	ListProducts(ctx context.Context, offset int, limit int) ([]*model.Product, error)
}
