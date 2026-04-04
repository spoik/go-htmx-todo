package createtodo

import (
	"net/http"

	"github.com/spoik/go-htmx-todo/internal/database/queries"
	"github.com/spoik/go-htmx-todo/internal/server/log"
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
		c.renderGenericError(w, r, title, err)
		return
	}

	todoVm, err := viewmodels.NewTodo(todo)

	if err != nil {
		c.renderGenericError(w, r, title, err)
		return
	}

	templates.Todo(todoVm).Render(r.Context(), w)
}

func (c CreateTodo) renderGenericError(w http.ResponseWriter, r *http.Request, title string, err error) {
	log.UnhandledError(r, err)
	err = templates.TodoForm(title, templates.GenericErrorMessage).Render(r.Context(), w)

	if err != nil {
		response.InternalServerError(w, r, err)
	}
}
