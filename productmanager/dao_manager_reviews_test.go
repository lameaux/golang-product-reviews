package productmanager

import (
	"context"
	"testing"

	"github.com/lameaux/golang-product-reviews/dto"
	"github.com/lameaux/golang-product-reviews/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDAOManager_CreateProductReview(t *testing.T) {
	dao := new(mockedDAO)
	dao.On("CreateProductReview", mock.Anything, &model.Review{
		ProductID: 2,
		FirstName: "Sergej",
		LastName:  "Sizov",
		Review:    "Excellent",
		Rating:    5,
	}).Return(1, nil)

	m := New(dao, func(ctx context.Context, productID model.ID, reviewID model.ID, action string) {
		assert.Equal(t, 2, productID)
		assert.Equal(t, 1, reviewID)
		assert.Equal(t, "create", action)
	})

	review := &dto.Review{
		FirstName: "Sergej",
		LastName:  "Sizov",
		Review:    "Excellent",
		Rating:    5,
	}

	id, err := m.CreateProductReview(t.Context(), 2, review)
	assert.NoError(t, err)
	assert.Equal(t, 1, id)
}

func TestDAOManager_UpdateProductReview(t *testing.T) {
	dao := new(mockedDAO)
	dao.On("UpdateProductReview", mock.Anything, &model.Review{
		ID:        1,
		ProductID: 2,
		FirstName: "Sergej",
		LastName:  "Sizov",
		Review:    "Excellent",
		Rating:    5,
	}).Return(nil)

	m := New(dao, func(ctx context.Context, productID model.ID, reviewID model.ID, action string) {
		assert.Equal(t, 2, productID)
		assert.Equal(t, 1, reviewID)
		assert.Equal(t, "update", action)
	})

	review := &dto.Review{
		FirstName: "Sergej",
		LastName:  "Sizov",
		Review:    "Excellent",
		Rating:    5,
	}

	err := m.UpdateProductReview(t.Context(), 2, 1, review)
	assert.NoError(t, err)
}

func TestDAOManager_DeleteProductReview(t *testing.T) {
	dao := new(mockedDAO)
	dao.On("DeleteProductReview", mock.Anything, 1).Return(nil)

	m := New(dao, func(ctx context.Context, productID model.ID, reviewID model.ID, action string) {
		assert.Equal(t, 2, productID)
		assert.Equal(t, 1, reviewID)
		assert.Equal(t, "delete", action)
	})

	err := m.DeleteProductReview(t.Context(), 2, 1)
	assert.NoError(t, err)
}

func TestDAOManager_GetProductReview(t *testing.T) {
	dao := new(mockedDAO)
	dao.On("GetProductReview", mock.Anything, 1).Return(&model.Review{
		ID:        1,
		ProductID: 2,
		FirstName: "Sergej",
		LastName:  "Sizov",
		Review:    "Excellent",
		Rating:    5,
	}, nil)

	m := New(dao, nil)

	product, err := m.GetProductReview(t.Context(), 2, 1)
	assert.NoError(t, err)

	assert.Equal(t, &dto.Review{
		ID:        1,
		FirstName: "Sergej",
		LastName:  "Sizov",
		Review:    "Excellent",
		Rating:    5,
	}, product)
}

func TestDAOManager_ListProductReviews(t *testing.T) {
	dao := new(mockedDAO)
	dao.On("ListProductReviews", mock.Anything, 2, 0, 100).Return(
		[]*model.Review{
			{
				ID:        1,
				ProductID: 2,
				FirstName: "Sergej",
				LastName:  "Sizov",
				Review:    "Excellent",
				Rating:    5,
			},
		}, nil)

	m := New(dao, nil)

	products, err := m.ListProductReviews(t.Context(), 2, 0, 100)
	assert.NoError(t, err)

	assert.Equal(t, []*dto.Review{
		{
			ID:        1,
			FirstName: "Sergej",
			LastName:  "Sizov",
			Review:    "Excellent",
			Rating:    5,
		},
	}, products)
}
