package templates

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/spoik/go-htmx-todo/internal/server/log"
	"github.com/spoik/go-htmx-todo/internal/server/response"
)

const GenericErrorMessage = "Something went wrong. Please try again."

func RenderOrInternalError(w http.ResponseWriter, r *http.Request, cmp templ.Component) {
	err := cmp.Render(r.Context(), w)

	if err != nil {
		response.InternalServerError(w, r, err)
	}
}

func GenericHTMLError(w http.ResponseWriter, r *http.Request, err error) {
	log.UnhandledError(r, err)
	RenderOrInternalError(w, r, Error(GenericErrorMessage))
}
