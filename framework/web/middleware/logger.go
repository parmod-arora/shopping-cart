package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"cinemo.com/shoping-cart/cmd/serverd/appenv"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type contextKey string

var loggerContextKey = contextKey("logger")

var instanceID = time.Now().Format("20060102150405") + "-" + appenv.GetWithDefault("CF_INSTANCE_INDEX", "X")

// RequestLogger serves as a middleware that logs the start and end of each request, along with some useful data as logger fields
func RequestLogger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ctx := r.Context()
		// Parse request information
		requestURIparts := append(strings.SplitN(r.RequestURI, "?", 2), "")

		// Instantiate verbose logger
		logger := logrus.
			WithField("request", uuid.New().String()).
			WithField("route", r.Method+" "+requestURIparts[0]).
			WithField("query", requestURIparts[1]).
			WithField("instance", instanceID).
			WithField("ip", r.RemoteAddr).
			WithField("referer", r.Referer()).
			WithField("agent", r.UserAgent())

		ctx = SetLogger(ctx, logger)
		logger.
			Infof("START")

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

		logger.
			WithField("duration", time.Since(start)).
			Infof("END")

	}
	return http.HandlerFunc(fn)
}

// SetLogger sets the logger into the provided context and returns a copy
func SetLogger(ctx context.Context, value *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggerContextKey, value)
}
