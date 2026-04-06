package templates

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/spoik/go-htmx-todo/internal/server/response"
)

func Render(w http.ResponseWriter, r *http.Request, cmp templ.Component) {
	err := cmp.Render(r.Context(), w)

	if err != nil {
		response.InternalServerError(w, r, err)
	}
}
