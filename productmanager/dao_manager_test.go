package productmanager

import (
	"context"

	"github.com/lameaux/golang-product-reviews/cache"
	"github.com/lameaux/golang-product-reviews/database"
	"github.com/lameaux/golang-product-reviews/lock"
	"github.com/lameaux/golang-product-reviews/model"
	"github.com/stretchr/testify/mock"
)

type mockedDAO struct {
	mock.Mock
}

type mockedCache struct {
	mock.Mock
}

type mockedLock struct {
	mock.Mock
}

var _ database.DAO = (*mockedDAO)(nil)
var _ cache.DAO = (*mockedCache)(nil)
var _ lock.Lock = (*mockedLock)(nil)

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
func (m *mockedDAO) DeleteProductReview(ctx context.Context, reviewID model.ID) error {
	args := m.Called(ctx, reviewID)
	return args.Error(0)
}
func (m *mockedDAO) GetProductReview(ctx context.Context, reviewID model.ID) (*model.Review, error) {
	args := m.Called(ctx, reviewID)
	return args.Get(0).(*model.Review), args.Error(1)
}
func (m *mockedDAO) ListProductReviews(ctx context.Context, productID model.ID, offset int, limit int) ([]*model.Review, error) {
	args := m.Called(ctx, productID, offset, limit)
	return args.Get(0).([]*model.Review), args.Error(1)
}

func (m *mockedCache) InvalidateProduct(ctx context.Context, productID model.ID) {
	m.Called(ctx, productID)
}

func (m *mockedCache) GetProductRating(ctx context.Context, productID model.ID) (float32, error) {
	args := m.Called(ctx, productID)
	return args.Get(0).(float32), args.Error(1)
}
func (m *mockedCache) SetProductRating(ctx context.Context, productID model.ID, rating float32) {
	m.Called(ctx, productID, rating)
}

func (m *mockedCache) GetProductReviews(ctx context.Context, productID model.ID, offset int, limit int) ([]*model.Review, error) {
	args := m.Called(ctx, productID, offset, limit)
	return args.Get(0).([]*model.Review), args.Error(1)
}

func (m *mockedCache) SetProductReviews(ctx context.Context, productID model.ID, offset int, limit int, reviews []*model.Review) {
	m.Called(ctx, productID, offset, limit, reviews)
}

func (m *mockedCache) GetProductReview(ctx context.Context, productID model.ID, reviewID model.ID) (*model.Review, error) {
	args := m.Called(ctx, productID, reviewID)
	return args.Get(0).(*model.Review), args.Error(1)
}

func (m *mockedCache) SetProductReview(ctx context.Context, productID model.ID, reviewID model.ID, review *model.Review) {
	m.Called(ctx, productID, reviewID, review)
}

func (m *mockedLock) Lock(ctx context.Context, productID model.ID) error {
	args := m.Called(ctx, productID)
	return args.Error(0)
}

func (m *mockedLock) Unlock(ctx context.Context, productID model.ID) error {
	args := m.Called(ctx, productID)
	return args.Error(0)
}
