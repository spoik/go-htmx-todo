package response

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/spoik/go-htmx-todo/internal/templates"
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

func GenericHTMLError(w http.ResponseWriter, r *http.Request, err error, component templ.Component) {
	templates.GenericError(templates.AddTodo()).Render(r.Context(), w)
	slog.Error(
		"Unhandled error",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err,
	)
}
