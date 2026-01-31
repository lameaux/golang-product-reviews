package lock

import (
	"context"

	"github.com/lameaux/golang-product-reviews/model"
)

type Lock interface {
	Lock(ctx context.Context, id model.ID) error
	Unlock(ctx context.Context, id model.ID) error
}
