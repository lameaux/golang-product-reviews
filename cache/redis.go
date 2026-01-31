package cache

import (
	"context"

	"github.com/lameaux/golang-product-reviews/model"
	"github.com/rs/zerolog"
)

var _ DAO = (*RedisCache)(nil)

type RedisCache struct {
	logger *zerolog.Logger
}

func NewRedis(logger *zerolog.Logger) *RedisCache {
	return &RedisCache{logger: logger}
}

func (r *RedisCache) InvalidateProduct(ctx context.Context, productID model.ID) {
	r.logger.Debug().Int("id", productID).Msg("InvalidateProduct")
}

func (r *RedisCache) GetProductRating(ctx context.Context, productID model.ID) (float32, error) {
	return 0, NotFound
}
func (r *RedisCache) SetProductRating(ctx context.Context, productID model.ID, rating float32) {
	r.logger.Debug().Int("id", productID).Float32("rating", rating).Msg("SetProductRating")
}

func (r *RedisCache) GetProductReview(ctx context.Context, productID model.ID, reviewID model.ID) (*model.Review, error) {
	r.logger.Debug().Int("id", reviewID).Msg("GetProductReview")
	return nil, NotFound
}
func (r *RedisCache) SetProductReview(ctx context.Context, productID model.ID, reviewID model.ID, review *model.Review) {
	r.logger.Debug().Int("id", reviewID).Msg("SetProductReview")
}

func (r *RedisCache) GetProductReviews(ctx context.Context, productID model.ID, offset int, limit int) ([]*model.Review, error) {
	r.logger.Debug().Int("id", productID).Msg("GetProductReviews")
	return nil, NotFound
}
func (r *RedisCache) SetProductReviews(ctx context.Context, productID model.ID, offset int, limit int, reviews []*model.Review) {
	r.logger.Debug().Int("id", productID).Msg("SetProductReviews")
}
