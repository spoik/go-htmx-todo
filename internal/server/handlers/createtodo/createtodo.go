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

type createTodo struct {
	queries *queries.Queries
}

func New(q *queries.Queries) createTodo {
	return createTodo{queries: q}
}

func (c createTodo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	todoListId := r.PostFormValue("todo_list_id")

	todoVm, err := c.insertTodo(r.Context(), todoListId, title)

	if err != nil {
		c.renderGenericError(w, r, todoListId, title, err)
		return
	}

	templates.Render(w, r, templates.TodoCreated(todoVm))
}

func (c createTodo) insertTodo(ctx context.Context, todoListId, title string) (todoVm viewmodels.Todo, err error) {
	var todoListUUID pgtype.UUID
	if err = todoListUUID.Scan(todoListId); err != nil {
		return
	}

	// TODO: Validate new todo

	todo, err := c.queries.InsertTodo(ctx, queries.InsertTodoParams{
		Title:      title,
		TodoListID: todoListUUID,
	})

	if err != nil {
		return
	}

	return viewmodels.NewTodo(todo)
}

func (c createTodo) renderGenericError(w http.ResponseWriter, r *http.Request, todoListId string, title string, err error) {
	log.UnhandledError(r, err)
	templates.Render(w, r, templates.NewTodoForm(todoListId, title, templates.GenericErrorMessage))
}
