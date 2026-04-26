package createtodo

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spoik/go-htmx-todo/internal/database/queries"
	"github.com/spoik/go-htmx-todo/internal/server/log"
	"github.com/spoik/go-htmx-todo/internal/templates"
	"github.com/spoik/go-htmx-todo/internal/templates/viewmodels"
)

type formTodo struct {
	todoListId string
	title      string
}

func newFormTodo(r *http.Request) formTodo {
	return formTodo{
		todoListId: r.PostFormValue("todo_list_id"),
		title:      r.PostFormValue("title"),
	}
}

type createTodo struct {
	queries *queries.Queries
}

func New(q *queries.Queries) createTodo {
	return createTodo{queries: q}
}

func (c createTodo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	formTodo := newFormTodo(r)

	// TODO: Validate new todo

	todoVm, err := c.insertTodo(r.Context(), formTodo)

	if err != nil {
		c.renderGenericError(w, r, formTodo, err)
		return
	}

	templates.Render(w, r, templates.TodoCreated(todoVm))
}

func (c createTodo) insertTodo(ctx context.Context, formTodo formTodo) (todoVm viewmodels.Todo, err error) {
	var todoListUUID pgtype.UUID
	if err = todoListUUID.Scan(formTodo.todoListId); err != nil {
		return
	}

	todo, err := c.queries.InsertTodo(ctx, queries.InsertTodoParams{
		Title:      formTodo.title,
		TodoListID: todoListUUID,
	})

	if err != nil {
		return
	}

	return viewmodels.NewTodo(todo)
}

func (c createTodo) renderGenericError(w http.ResponseWriter, r *http.Request, formTodo formTodo, err error) {
	log.UnhandledError(r, err)

	template := templates.NewTodoForm(
		formTodo.todoListId,
		formTodo.title,
		templates.GenericErrorMessage,
	)

	templates.Render(w, r, template)
}
