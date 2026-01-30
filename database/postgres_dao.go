package database

import (
	"context"

	"github.com/lameaux/golang-product-reviews/model"
	"gorm.io/gorm"
)

var _ DAO = (*postgresDAO)(nil)

type postgresDAO struct {
	db *gorm.DB
}

func NewPostgresDAO(db *gorm.DB) *postgresDAO {
	return &postgresDAO{
		db: db,
	}
}

func (d *postgresDAO) CreateProduct(ctx context.Context, p *model.Product) (model.ID, error) {
	return 0, nil
}

func (d *postgresDAO) UpdateProduct(ctx context.Context, p *model.Product) error {
	return nil
}

func (d *postgresDAO) DeleteProduct(ctx context.Context, id model.ID) error {
	return nil
}

func (d *postgresDAO) GetProduct(ctx context.Context, id model.ID) (*model.Product, error) {
	return nil, nil
}

func (d *postgresDAO) GetProductRating(ctx context.Context, id model.ID) (float32, error) {
	return 0.0, nil
}

func (d *postgresDAO) ListProducts(ctx context.Context, offset int, limit int) ([]*model.Product, error) {
	return nil, nil
}
