package notifier

import (
	"fmt"

	"github.com/lameaux/golang-product-reviews/model"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
)

type Notifier struct {
	logger   *zerolog.Logger
	natsConn *nats.Conn
}

func New(logger *zerolog.Logger, natsConn *nats.Conn) *Notifier {
	return &Notifier{logger: logger, natsConn: natsConn}
}

func (n *Notifier) Notify(productID model.ID, reviewID model.ID, action string) {
	msg := fmt.Sprintf(
		`{"product":%d,"review":%d,"action":"%s"}`,
		productID, reviewID, action,
	)

	n.logger.Info().
		Str("msg", msg).
		Msg("notify")

	if err := n.natsConn.Publish("reviews", []byte(msg)); err != nil {
		n.logger.Error().Err(err).Msg("field to publish message to NATS")
	}
}
