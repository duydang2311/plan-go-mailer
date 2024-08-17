package mailing

import (
	"github.com/resend/resend-go/v2"
)

type ResendOptions struct {
	ApiKey string
}

func NewResend(options *ResendOptions) *resend.Client {
	return resend.NewClient(options.ApiKey)
}
