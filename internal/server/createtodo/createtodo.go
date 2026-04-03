package createtodo

import (
	"net/http"

	"github.com/spoik/go-htmx-todo/internal/database/queries"
	"github.com/spoik/go-htmx-todo/internal/server/response"
	"github.com/spoik/go-htmx-todo/internal/templates"
	"github.com/spoik/go-htmx-todo/internal/templates/viewmodels"
)

type CreateTodo struct {
	queries *queries.Queries
}

func New(q *queries.Queries) CreateTodo {
	return CreateTodo{queries: q}
}

func (c CreateTodo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")

	// TODO: Validate new todo
	// templates.TodoForm(title, title).Render(r.Context(), w)

	todo, err := c.queries.InsertTodo(r.Context(), title)

	if err != nil {
		response.InternalServerError(w, r, err)
	}

	todoVm, err := viewmodels.NewTodo(todo)

	if err != nil {
		response.InternalServerError(w, r, err)
	}

	templates.Todo(todoVm).Render(r.Context(), w)
}
