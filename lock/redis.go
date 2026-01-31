package lock

import (
	"context"

	"github.com/lameaux/golang-product-reviews/model"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

var _ Lock = (*RedisLock)(nil)

type RedisLock struct {
	logger *zerolog.Logger
	client *redis.Client
}

func NewRedis(logger *zerolog.Logger, client *redis.Client) *RedisLock {
	return &RedisLock{logger: logger, client: client}
}

func (r *RedisLock) Lock(ctx context.Context, id model.ID) error {
	r.logger.Debug().Int("id", id).Msg("redis lock")
	return nil
}

func (r *RedisLock) Unlock(ctx context.Context, id model.ID) error {
	r.logger.Debug().Int("id", id).Msg("redis unlock")
	return nil
}
