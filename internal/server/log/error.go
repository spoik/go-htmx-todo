package log

import (
	"log/slog"
	"net/http"
)

// Gives a convenient way to log errors that can't be recovered from elegantly.
func UnhandledError(r *http.Request, err error) {
	slog.Error(
		"Unhandled error",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err,
	)
}
