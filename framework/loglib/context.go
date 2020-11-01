package loglib

import (
	"context"

	"github.com/sirupsen/logrus"
)

type contextKey string

var loggerContextKey = contextKey("logger")

// SetLogger sets the logger into the provided context and returns a copy
func SetLogger(ctx context.Context, value *logrus.Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey, value)
}

// GetLogger returns logger object from Context, else, return default concise logger
func GetLogger(ctx context.Context) *logrus.Logger {
	if logger, ok := ctx.Value(loggerContextKey).(*logrus.Logger); ok {
		return logger
	}
	return logrus.StandardLogger()
}
