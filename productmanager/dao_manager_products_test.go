package productmanager

import (
	"testing"

	"github.com/lameaux/golang-product-reviews/cache"
	"github.com/lameaux/golang-product-reviews/dto"
	"github.com/lameaux/golang-product-reviews/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDAOManager_CreateProduct(t *testing.T) {
	dao := new(mockedDAO)
	dao.On("CreateProduct", mock.Anything, &model.Product{
		Name:        "P1",
		Description: "P1 desc",
		Price:       100,
	}).Return(1, nil)

	m := New(dao, nil, nil, nil)

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

	m := New(dao, nil, nil, nil)

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

	cacheDAO := new(mockedCache)
	cacheDAO.On("InvalidateProduct", mock.Anything, 1).Once()

	m := New(dao, cacheDAO, nil, func(productID model.ID, reviewID model.ID, action string) {
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

	cacheDAO := new(mockedCache)
	cacheDAO.On("GetProductRating", mock.Anything, 1).Return(float32(0), cache.NotFound).Twice()
	cacheDAO.On("SetProductRating", mock.Anything, 1, float32(4.9)).Once()

	lock := new(mockedLock)
	lock.On("Lock", mock.Anything, 1).Return(nil)
	lock.On("Unlock", mock.Anything, 1).Return(nil)

	m := New(dao, cacheDAO, lock, nil)

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

	cacheDAO := new(mockedCache)
	cacheDAO.On("GetProductRating", mock.Anything, 1).Return(float32(0), cache.NotFound).Twice()
	cacheDAO.On("SetProductRating", mock.Anything, 1, float32(4.9)).Once()

	lock := new(mockedLock)
	lock.On("Lock", mock.Anything, 1).Return(nil)
	lock.On("Unlock", mock.Anything, 1).Return(nil)

	m := New(dao, cacheDAO, lock, nil)

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
