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

	if product == nil {
		return nil, nil
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
	review := &model.Review{
		ProductID: productID,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Review:    r.Review,
		Rating:    r.Rating,
	}

	reviewID, err := m.dao.CreateProductReview(ctx, review)
	if err != nil {
		return 0, fmt.Errorf("dao.CreateProductReview: %w", err)
	}

	// TODO: save to cache

	m.notifyFunc(ctx, productID, reviewID, "create")

	return reviewID, nil
}

func (m *DAOManager) UpdateProductReview(ctx context.Context, productID model.ID, reviewID model.ID, r *dto.Review) error {
	review := &model.Review{
		ID:        reviewID,
		ProductID: productID,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Review:    r.Review,
		Rating:    r.Rating,
	}

	if err := m.dao.UpdateProductReview(ctx, review); err != nil {
		return fmt.Errorf("dao.UpdateProductReview: %w", err)
	}

	// TODO: save to cache

	m.notifyFunc(ctx, productID, reviewID, "update")

	return nil
}

func (m *DAOManager) DeleteProductReview(ctx context.Context, productID model.ID, reviewID model.ID) error {
	if err := m.dao.DeleteProductReview(ctx, productID, reviewID); err != nil {
		return fmt.Errorf("dao.DeleteProductReview: %w", err)
	}

	// TODO: invalidate cache

	m.notifyFunc(ctx, productID, reviewID, "delete")

	return nil
}

func (m *DAOManager) GetProductReview(ctx context.Context, productID model.ID, reviewID model.ID) (*dto.Review, error) {
	// TODO: load reviews from cache

	review, err := m.dao.GetProductReview(ctx, productID, reviewID)
	if err != nil {
		return nil, fmt.Errorf("dao.GetProductReview: %w", err)
	}

	if review == nil {
		return nil, nil
	}

	return convertReview(review), nil
}

func (m *DAOManager) ListProductReviews(ctx context.Context, productID model.ID, offset int, limit int) ([]*dto.Review, error) {
	// TODO: load reviews from cache

	reviews, err := m.dao.ListProductReviews(ctx, productID, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("dao.ListProductReviews: %w", err)
	}

	result := make([]*dto.Review, 0, len(reviews))
	for _, product := range reviews {
		result = append(result, convertReview(product))
	}

	return result, nil
}

func convertReview(review *model.Review) *dto.Review {
	return &dto.Review{
		ID:        review.ID,
		FirstName: review.FirstName,
		LastName:  review.LastName,
		Review:    review.Review,
		Rating:    review.Rating,
	}
}
