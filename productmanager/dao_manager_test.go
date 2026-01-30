package productmanager

import (
	"context"
	"testing"

	"github.com/lameaux/golang-product-reviews/database"
	"github.com/lameaux/golang-product-reviews/dto"
	"github.com/lameaux/golang-product-reviews/model"
	"github.com/stretchr/testify/assert"
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
func TestDAOManager_CreateProduct(t *testing.T) {
	dao := new(mockedDAO)
	dao.On("CreateProduct", mock.Anything, &model.Product{
		Name:        "P1",
		Description: "P1 desc",
		Price:       100,
	}).Return(1, nil)

	m := New(dao, nil)

	p := &dto.Product{
		Name:        "P1",
		Description: "P1 desc",
		Price:       100,
	}

	id, err := m.CreateProduct(t.Context(), p)
	assert.NoError(t, err)
	assert.Equal(t, 1, id)
}

func TestDAOManager_UpdateProduct(t *testing.T) {
	dao := new(mockedDAO)
	dao.On("UpdateProduct", mock.Anything, &model.Product{
		ID:          1,
		Name:        "P1",
		Description: "P1 desc",
		Price:       100,
	}).Return(nil)

	m := New(dao, nil)

	p := &dto.Product{
		Name:        "P1",
		Description: "P1 desc",
		Price:       100,
	}

	err := m.UpdateProduct(t.Context(), 1, p)
	assert.NoError(t, err)
}

func TestDAOManager_DeleteProduct(t *testing.T) {
	dao := new(mockedDAO)
	dao.On("DeleteProduct", mock.Anything, 1).Return(nil)

	m := New(dao, func(ctx context.Context, productID model.ID, reviewID model.ID, action string) {
		assert.Equal(t, 1, productID)
		assert.Equal(t, 0, reviewID)
		assert.Equal(t, "delete", action)
	})

	err := m.DeleteProduct(t.Context(), 1)
	assert.NoError(t, err)
}

func TestDAOManager_GetProduct(t *testing.T) {
	dao := new(mockedDAO)
	dao.On("GetProduct", mock.Anything, 1).Return(&model.Product{
		ID:          1,
		Name:        "P1",
		Description: "P1 desc",
		Price:       100,
	}, nil)

	dao.On("GetProductRating", mock.Anything, 1).Return(float32(4.9), nil)

	m := New(dao, nil)

	product, err := m.GetProduct(t.Context(), 1)
	assert.NoError(t, err)

	assert.Equal(t, &dto.ProductWithRating{
		Product: dto.Product{
			ID:          1,
			Name:        "P1",
			Description: "P1 desc",
			Price:       100,
		},
		Rating: 4.9,
	}, product)
}

func TestDAOManager_ListProducts(t *testing.T) {
	dao := new(mockedDAO)
	dao.On("ListProducts", mock.Anything, 0, 100).Return(
		[]*model.Product{
			{
				ID:          1,
				Name:        "P1",
				Description: "P1 desc",
				Price:       100,
			},
		}, nil)

	dao.On("GetProductRating", mock.Anything, 1).Return(float32(4.9), nil)

	m := New(dao, nil)

	products, err := m.ListProducts(t.Context(), 0, 100)
	assert.NoError(t, err)

	assert.Equal(t, []*dto.ProductWithRating{
		{
			Product: dto.Product{
				ID:          1,
				Name:        "P1",
				Description: "P1 desc",
				Price:       100,
			},
			Rating: 4.9,
		},
	}, products)
}
