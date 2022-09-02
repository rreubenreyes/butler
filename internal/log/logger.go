package log

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

type logger struct{}

var wr = &zerolog.ConsoleWriter{
	Out:        os.Stdout,
	PartsOrder: []string{"time", "level", "message"},
}

// ForContext copies the provided context and adds a *zerolog.Logger.
func ForContext(ctx context.Context) context.Context {
	l := zerolog.New(wr).With().Timestamp().Logger()

	return context.WithValue(ctx, logger{}, &l)
}

// ForContext copies the provided context and adds a *zerolog.Logger.
func ForContextWith(ctx context.Context, key string, value string) context.Context {
	l := zerolog.New(wr).With().Timestamp().Str(key, value).Logger()

	return context.WithValue(ctx, logger{}, &l)
}

// FromContext gets a *zerolog.Logger from the provided context object.
func FromContext(ctx context.Context) *zerolog.Logger {
	if l, ok := ctx.Value(logger{}).(*zerolog.Logger); ok {
		return l
	}

	l := zerolog.New(wr).With().Timestamp().Logger()

	return &l
}
