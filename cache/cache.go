package cache

import (
	"context"
	"errors"

	"github.com/lameaux/golang-product-reviews/model"
)

var NotFound = errors.New("not found")

type DAO interface {
	InvalidateProduct(ctx context.Context, id model.ID)

	GetProductRating(ctx context.Context, productID model.ID) (float32, error)
	SetProductRating(ctx context.Context, productID model.ID, rating float32)

	GetProductReview(ctx context.Context, productID model.ID, reviewID model.ID) (*model.Review, error)
	SetProductReview(ctx context.Context, productID model.ID, reviewID model.ID, review *model.Review)

	GetProductReviews(ctx context.Context, productID model.ID, offset int, limit int) ([]*model.Review, error)
	SetProductReviews(ctx context.Context, productID model.ID, offset int, limit int, reviews []*model.Review)
}
