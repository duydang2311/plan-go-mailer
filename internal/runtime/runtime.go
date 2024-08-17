package runtime

import (
	"context"

	"github.com/nats-io/nats.go"
	"github.com/resend/resend-go/v2"
)

type Runtime struct {
	Context *context.Context
	Nats    *nats.Conn
	Mailer  *resend.Client
}
