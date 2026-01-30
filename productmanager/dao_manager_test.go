package productmanager

import (
	"context"

	"github.com/lameaux/golang-product-reviews/database"
	"github.com/lameaux/golang-product-reviews/model"
	"github.com/stretchr/testify/mock"
)

type mockedDAO struct {
	mock.Mock
}

var _ database.DAO = (*mockedDAO)(nil)

func (m *mockedDAO) CreateProduct(ctx context.Context, product *model.Product) (model.ID, error) {
	args := m.Called(ctx, product)
	return args.Int(0), args.Error(1)
}

func (m *mockedDAO) UpdateProduct(ctx context.Context, product *model.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *mockedDAO) DeleteProduct(ctx context.Context, id model.ID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *mockedDAO) GetProductRating(ctx context.Context, id model.ID) (float32, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(float32), args.Error(1)
}

func (m *mockedDAO) GetProduct(ctx context.Context, id model.ID) (*model.Product, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Product), args.Error(1)
}

func (m *mockedDAO) ListProducts(ctx context.Context, offset int, limit int) ([]*model.Product, error) {
	args := m.Called(ctx, offset, limit)
	return args.Get(0).([]*model.Product), args.Error(1)
}

func (m *mockedDAO) CreateProductReview(ctx context.Context, review *model.Review) (model.ID, error) {
	args := m.Called(ctx, review)
	return args.Int(0), args.Error(1)
}
func (m *mockedDAO) UpdateProductReview(ctx context.Context, review *model.Review) error {
	args := m.Called(ctx, review)
	return args.Error(0)
}
func (m *mockedDAO) DeleteProductReview(ctx context.Context, productID model.ID, reviewID model.ID) error {
	args := m.Called(ctx, productID, reviewID)
	return args.Error(0)
}
func (m *mockedDAO) GetProductReview(ctx context.Context, productID model.ID, reviewID model.ID) (*model.Review, error) {
	args := m.Called(ctx, productID, reviewID)
	return args.Get(0).(*model.Review), args.Error(1)
}
func (m *mockedDAO) ListProductReviews(ctx context.Context, productID model.ID, offset int, limit int) ([]*model.Review, error) {
	args := m.Called(ctx, productID, offset, limit)
	return args.Get(0).([]*model.Review), args.Error(1)
}
