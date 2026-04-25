package createtodolist

import (
	"net/http"

	"github.com/spoik/go-htmx-todo/internal/database/queries"
	"github.com/spoik/go-htmx-todo/internal/server/log"
	"github.com/spoik/go-htmx-todo/internal/templates"
	"github.com/spoik/go-htmx-todo/internal/templates/viewmodels"
)

type CreateTodoList struct {
	queries *queries.Queries
}

func New(q *queries.Queries) CreateTodoList {
	return CreateTodoList{queries: q}
}

func (c CreateTodoList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")

	if title == "" {
		templates.Render(w, r, templates.TodoListForm(title, "Title is required"))
		return
	}

	tl, err := c.queries.InsertTodoList(r.Context(), title)
	if err != nil {
		log.UnhandledError(r, err)
		templates.Render(w, r, templates.TodoListForm(title, templates.GenericErrorMessage))
		return
	}

	tlVm, err := viewmodels.NewTodoList(tl)
	if err != nil {
		log.UnhandledError(r, err)
		templates.Render(w, r, templates.TodoListForm(title, templates.GenericErrorMessage))
		return
	}

	atlVm, err := viewmodels.NewAddTodoList()
	if err != nil {
		log.UnhandledError(r, err)
		templates.Render(w, r, templates.TodoListForm(title, templates.GenericErrorMessage))
		return
	}

	templates.Render(w, r, templates.TodoListCreated(tlVm, atlVm))
}
