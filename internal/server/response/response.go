package response

import (
	"log/slog"
	"net/http"
)

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	slog.Error(
		"Internal server error",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err,
	)
}
