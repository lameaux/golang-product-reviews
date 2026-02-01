package lock

import (
	"context"
	"errors"

	"github.com/lameaux/golang-product-reviews/model"
)

var ErrLocked = errors.New("locked")

type Lock interface {
	Lock(ctx context.Context, id model.ID) error
	Unlock(ctx context.Context, id model.ID) error
}
