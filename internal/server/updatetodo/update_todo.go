package updatetodo

import (
	"net/http"
	"strconv"

	"github.com/spoik/go-htmx-todo/internal/database/queries"
	"github.com/spoik/go-htmx-todo/internal/server/response"
	"github.com/spoik/go-htmx-todo/internal/templates"
)

type updateTodo struct {
	queries *queries.Queries
}

func New(q *queries.Queries) updateTodo {
	return updateTodo{queries: q}
}

func (u updateTodo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	idInt, err := strconv.ParseInt(id, 10, 32)

	if err != nil {
		http.Error(
			w,
			"Invalid todo id. Must be an integer.",
			http.StatusUnprocessableEntity,
		)
		return
	}

	todo, err := u.queries.GetTodo(r.Context(), int32(idInt))

	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	todo.Complete.Bool = !todo.Complete.Bool

	templates.Todo(todo).Render(r.Context(), w)
}
