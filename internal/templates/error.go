package templates

import (
	"net/http"

	"github.com/spoik/go-htmx-todo/internal/server/log"
)

const GenericErrorMessage = "Something went wrong. Please try again."

// Convenient way to log and show a generic HTML error for errors that can't be recovered from elegantly.
func UnhandledError(w http.ResponseWriter, r *http.Request, err error) {
	log.UnhandledError(r, err)
	GenericHTMLError(w, r)
}

func GenericHTMLError(w http.ResponseWriter, r *http.Request) {
	Render(w, r, Error(GenericErrorMessage))
}

