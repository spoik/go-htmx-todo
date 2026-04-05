package deletetodo

import (
	"net/http"
	"strconv"

	"github.com/spoik/go-htmx-todo/internal/database/queries"
	"github.com/spoik/go-htmx-todo/internal/templates"
)

type deleteTodo struct {
	queries *queries.Queries
}

func New(q *queries.Queries) deleteTodo {
	return deleteTodo{queries: q}
}

func (d deleteTodo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	parsedId, err := strconv.ParseInt(id, 10, 32)

	idInt := int32(parsedId)

	if err != nil {
		templates.GenericHTMLError(w, r)
		return
	}

	exists, err := d.queries.TodoExists(r.Context(), idInt)

	if err != nil {
		templates.UnhandledError(w, r, err)
		return
	}

	if !exists {
		w.Write([]byte{})
		return
	}

	err = d.queries.DeleteTodo(r.Context(), idInt)

	if err != nil {
		templates.UnhandledError(w, r, err)
		return
	}

	w.Write([]byte{})
	return
}
