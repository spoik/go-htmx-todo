package response

import (
	"log/slog"
	"net/http"
)

func InternalServerError(w http.ResponseWriter, err error) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	slog.Error("Internal server error", "error", err)
}
