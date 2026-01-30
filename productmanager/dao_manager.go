package productmanager

import (
	"context"
	"fmt"

	"github.com/lameaux/golang-product-reviews/database"
	"github.com/lameaux/golang-product-reviews/dto"
	"github.com/lameaux/golang-product-reviews/model"
)

var _ Manager = (*DAOManager)(nil)

type NotifyFunc func(ctx context.Context, productID model.ID, reviewID model.ID, action string)

type DAOManager struct {
	dao        database.DAO
	notifyFunc NotifyFunc
}

func New(
	dao database.DAO,
	notifyFunc NotifyFunc,
) *DAOManager {
	return &DAOManager{dao: dao, notifyFunc: notifyFunc}
}

func (m *DAOManager) CreateProduct(ctx context.Context, p *dto.Product) (model.ID, error) {
	product := &model.Product{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}

	productID, err := m.dao.CreateProduct(ctx, product)
	if err != nil {
		return 0, fmt.Errorf("dao.CreateProduct: %w", err)
	}

	return productID, nil
}

func (m *DAOManager) UpdateProduct(ctx context.Context, productID model.ID, p *dto.Product) error {
	product := &model.Product{
		ID:          productID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}

	if err := m.dao.UpdateProduct(ctx, product); err != nil {
		return fmt.Errorf("dao.UpdateProduct: %w", err)
	}

	return nil
}

func (m *DAOManager) DeleteProduct(ctx context.Context, productID model.ID) error {
	if err := m.dao.DeleteProduct(ctx, productID); err != nil {
		return fmt.Errorf("dao.UpdateProduct: %w", err)
	}

	m.notifyFunc(ctx, productID, 0, "delete")

	return nil
}

func (m *DAOManager) GetProduct(ctx context.Context, productID model.ID) (*dto.ProductWithRating, error) {
	product, err := m.dao.GetProduct(ctx, productID)
	if err != nil {
		return nil, fmt.Errorf("dao.GetProduct: %w", err)
	}

	rating, err := m.getProductRating(ctx, product.ID)
	if err != nil {
		return nil, fmt.Errorf("getProductRating: %w", err)
	}

	return convertProductWithRating(product, rating), nil
}

func (m *DAOManager) ListProducts(ctx context.Context, offset int, limit int) ([]*dto.ProductWithRating, error) {
	products, err := m.dao.ListProducts(ctx, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("dao.ListProducts: %w", err)
	}

	result := make([]*dto.ProductWithRating, 0, len(products))
	for _, product := range products {
		rating, err := m.getProductRating(ctx, product.ID)
		if err != nil {
			return nil, fmt.Errorf("getProductRating: %w", err)
		}

		result = append(result, convertProductWithRating(product, rating))
	}

	return result, nil
}

func convertProductWithRating(product *model.Product, rating float32) *dto.ProductWithRating {
	return &dto.ProductWithRating{
		Product: dto.Product{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		},
		Rating: rating,
	}
}

func (m *DAOManager) getProductRating(ctx context.Context, id model.ID) (float32, error) {
	// TODO: load from cache

	// fallback to DB
	rating, err := m.dao.GetProductRating(ctx, id)
	if err != nil {
		return 0, fmt.Errorf("dao.GetProductAverageRating: %w", err)
	}

	// TODO: save to cache

	return rating, nil
}

func (m *DAOManager) CreateProductReview(ctx context.Context, productID model.ID, r *dto.Review) (model.ID, error) {
	reviewID := 0
	m.notifyFunc(ctx, productID, reviewID, "create")
	return 0, nil
}

func (m *DAOManager) UpdateProductReview(ctx context.Context, productID model.ID, reviewID model.ID, review *dto.Review) error {
	m.notifyFunc(ctx, productID, reviewID, "update")
	return nil
}

func (m *DAOManager) DeleteProductReview(ctx context.Context, productID model.ID, reviewID model.ID) error {
	m.notifyFunc(ctx, productID, reviewID, "delete")
	return nil
}

func (*DAOManager) GetProductReview(ctx context.Context, productID model.ID, reviewID model.ID) (*dto.Review, error) {
	return nil, nil
}

func (*DAOManager) ListProductReviews(ctx context.Context, productID model.ID, offset int, limit int) ([]*dto.Review, error) {
	return []*dto.Review{}, nil
}
