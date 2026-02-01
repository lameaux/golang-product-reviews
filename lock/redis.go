package lock

import (
	"context"
	"fmt"
	"time"

	"github.com/lameaux/golang-product-reviews/model"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

var _ Lock = (*RedisLock)(nil)

const ttl = 10 * time.Second
const prefix = "products:locks"

const maxRetry = 5

type RedisLock struct {
	logger *zerolog.Logger
	client *redis.Client
}

func NewRedis(logger *zerolog.Logger, client *redis.Client) *RedisLock {
	return &RedisLock{logger: logger, client: client}
}

func (r *RedisLock) Lock(ctx context.Context, id model.ID) error {
	key := fmt.Sprintf("%s:%d", prefix, id)

	for i := 0; i < maxRetry; i++ {
		ok, err := r.client.SetNX(ctx, key, "1", ttl).Result()
		if err != nil {
			return fmt.Errorf("lock: %w", err)
		}

		if ok {
			r.logger.Debug().Str("key", key).Msg("redis lock")
			return nil
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Duration(i) * time.Second):
		}
	}

	return ErrLocked
}

func (r *RedisLock) Unlock(ctx context.Context, id model.ID) error {
	key := fmt.Sprintf("%s:%d", prefix, id)

	if err := r.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("unlock: %w", err)
	}

	r.logger.Debug().Str("key", key).Msg("redis unlock")
	return nil
}
