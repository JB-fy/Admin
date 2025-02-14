package model

import (
	"context"
)

type EmailFunc func(ctx context.Context, config map[string]any) Email

type Email interface {
	GetFromEmail() (fromEmail string)
	SendEmail(ctx context.Context, message string, toEmailArr ...string) (err error)
}
