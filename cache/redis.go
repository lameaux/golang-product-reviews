package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/lameaux/golang-product-reviews/model"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

var _ DAO = (*RedisCache)(nil)

const prefix = "products"
const ttl = 1 * time.Hour

type RedisCache struct {
	logger *zerolog.Logger
	client *redis.Client
}

func NewRedis(logger *zerolog.Logger, client *redis.Client) *RedisCache {
	return &RedisCache{logger: logger, client: client}
}

func (r *RedisCache) InvalidateProduct(ctx context.Context, productID model.ID) {
	pattern := fmt.Sprintf("%s:%d:*", prefix, productID)

	var cursor uint64
	for {
		keys, next, err := r.client.Scan(ctx, cursor, pattern, 100).Result()
		if err != nil {
			r.logger.Warn().Err(err).Str("pattern", pattern).
				Msg("InvalidateProduct Scan failed")
			return
		}
		if len(keys) > 0 {
			if err := r.client.Del(ctx, keys...).Err(); err != nil {
				r.logger.Warn().Err(err).Str("pattern", pattern).
					Msg("InvalidateProduct Scan failed")
				return
			}
		}
		cursor = next
		if cursor == 0 {
			break
		}
	}

	r.logger.Debug().Str("pattern", pattern).Msg("InvalidateProduct")
}

func (r *RedisCache) GetProductRating(ctx context.Context, productID model.ID) (float32, error) {
	key := fmt.Sprintf("%s:%d:rating", prefix, productID)

	rating, err := r.client.Get(ctx, key).Float32()
	if err == redis.Nil {
		r.logger.Debug().Str("key", key).Msg("GetProductRating not found")
		return 0, NotFound
	}
	if err != nil {
		return 0, fmt.Errorf("GetProductRating: %w", err)
	}

	r.logger.Debug().Str("key", key).Float32("rating", rating).Msg("GetProductRating")
	return rating, nil
}
func (r *RedisCache) SetProductRating(ctx context.Context, productID model.ID, rating float32) {
	key := fmt.Sprintf("%s:%d:rating", prefix, productID)

	if err := r.client.Set(ctx, key, rating, ttl).Err(); err != nil {
		r.logger.Warn().Err(err).Str("key", key).Float32("rating", rating).Msg("SetProductRating failed")
		return
	}

	r.logger.Debug().Str("key", key).Float32("rating", rating).Msg("SetProductRating")
}

func (r *RedisCache) GetProductReview(ctx context.Context, productID model.ID, reviewID model.ID) (*model.Review, error) {
	key := fmt.Sprintf("%s:%d:review:%d", prefix, productID, reviewID)

	bytes, err := r.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		r.logger.Debug().Str("key", key).Msg("GetProductReview not found")
		return nil, NotFound
	}
	if err != nil {
		return nil, fmt.Errorf("GetProductReview: %w", err)
	}

	var review model.Review
	if err = json.Unmarshal(bytes, &review); err != nil {
		return nil, fmt.Errorf("GetProductReview unmarshal: %w", err)
	}

	r.logger.Debug().Str("key", key).Msg("GetProductReview")

	return &review, nil
}
func (r *RedisCache) SetProductReview(ctx context.Context, productID model.ID, reviewID model.ID, review *model.Review) {
	key := fmt.Sprintf("%s:%d:review:%d", prefix, productID, reviewID)

	bytes, err := json.Marshal(review)
	if err != nil {
		r.logger.Warn().Err(err).Str("key", key).Msg("SetProductReview marshal failed")
		return
	}

	if err := r.client.Set(ctx, key, bytes, ttl).Err(); err != nil {
		r.logger.Warn().Err(err).Str("key", key).Msg("SetProductReview failed")
		return
	}

	r.logger.Debug().Str("key", key).Msg("SetProductReview")
}

func (r *RedisCache) GetProductReviews(ctx context.Context, productID model.ID, offset int, limit int) ([]*model.Review, error) {
	key := fmt.Sprintf("%s:%d:reviews:%d:%d", prefix, productID, offset, limit)

	bytes, err := r.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		r.logger.Debug().Str("key", key).Msg("GetProductReviews not found")
		return nil, NotFound
	}
	if err != nil {
		return nil, fmt.Errorf("GetProductReviews: %w", err)
	}

	var reviews []*model.Review
	if err = json.Unmarshal(bytes, &reviews); err != nil {
		return nil, fmt.Errorf("GetProductReviews unmarshal: %w", err)
	}

	r.logger.Debug().Str(key, key).Msg("GetProductReviews")

	return reviews, nil
}
func (r *RedisCache) SetProductReviews(ctx context.Context, productID model.ID, offset int, limit int, reviews []*model.Review) {
	key := fmt.Sprintf("%s:%d:reviews:%d:%d", prefix, productID, offset, limit)

	bytes, err := json.Marshal(reviews)
	if err != nil {
		r.logger.Warn().Err(err).Str("key", key).Msg("SetProductReviews marshal failed")
		return
	}

	if err := r.client.Set(ctx, key, bytes, ttl).Err(); err != nil {
		r.logger.Warn().Err(err).Str("key", key).Msg("SetProductReviews failed")
		return
	}

	r.logger.Debug().Str(key, key).Msg("SetProductReviews")
}
