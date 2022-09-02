package log

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

type logger struct{}

// ForContext copies the provided context and adds a *zerolog.Logger.
func ForContext(ctx context.Context) context.Context {
	l := zerolog.New(os.Stdout)

	return context.WithValue(ctx, logger{}, &l)
}

// FromContext gets a *zerolog.Logger from the provided context object.
func FromContext(ctx context.Context) *zerolog.Logger {
	if l, ok := ctx.Value(logger{}).(*zerolog.Logger); ok {
		return l
	}

	l := zerolog.New(os.Stdout)

	return &l
}
