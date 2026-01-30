package notifier

import (
	"context"

	"github.com/lameaux/golang-product-reviews/model"
	"github.com/rs/zerolog"
)

type Notifier struct {
	logger *zerolog.Logger
}

func New(logger *zerolog.Logger) *Notifier {
	return &Notifier{logger: logger}
}

func (n *Notifier) Notify(ctx context.Context, productID model.ID, reviewID model.ID, action string) {
	n.logger.Info().
		Int("product", productID).
		Int("review", reviewID).
		Str("action", action).
		Msg("notify")
}
