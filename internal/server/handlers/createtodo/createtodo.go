package createtodo

import (
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
	todoListIdStr := r.PostFormValue("todo_list_id")

	var todoListId pgtype.UUID
	err := todoListId.Scan(todoListIdStr)
	if err != nil {
		c.renderGenericError(w, r, todoListIdStr, title, err)
		return
	}

	// TODO: Validate new todo

	todo, err := c.queries.InsertTodo(r.Context(), queries.InsertTodoParams{
		Title:      title,
		TodoListID: todoListId,
	})

	if err != nil {
		c.renderGenericError(w, r, todoListIdStr, title, err)
		return
	}

	todoVm, err := viewmodels.NewTodo(todo)

	if err != nil {
		c.renderGenericError(w, r, todoListIdStr, title, err)
		return
	}

	templates.Render(w, r, templates.TodoCreated(todoVm))
}

func (c createTodo) renderGenericError(w http.ResponseWriter, r *http.Request, todoListId string, title string, err error) {
	log.UnhandledError(r, err)
	templates.Render(w, r, templates.NewTodoForm(todoListId, title, templates.GenericErrorMessage))
}
