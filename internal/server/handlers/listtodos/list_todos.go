package listtodos

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spoik/go-htmx-todo/internal/database/queries"
	"github.com/spoik/go-htmx-todo/internal/server/response"
	"github.com/spoik/go-htmx-todo/internal/templates"
	"github.com/spoik/go-htmx-todo/internal/templates/viewmodels"
)

type listTodos struct {
	queries *queries.Queries
}

func New(q *queries.Queries) listTodos {
	return listTodos{queries: q}
}

func (l listTodos) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	todoListId := r.PathValue("todoListId")
	if todoListId == "" {
		http.NotFound(w, r)
		return
	}

	todos, err := l.getTodos(r.Context(), todoListId)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	todoLists, err := l.getTodoLists(r.Context())
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	templates.Render(w, r, templates.SidebarAndTodos(todoListId, todos, todoLists))
}

func (l listTodos) getTodos(c context.Context, todoListId string) (vms []viewmodels.Todo, err error) {
	var todoListUUID pgtype.UUID
	if err = todoListUUID.Scan(todoListId); err != nil {
		return
	}

	todos, err := l.queries.GetTodos(c, todoListUUID)
	if err != nil {
		return
	}

	return viewmodels.NewTodos(todos)
}

func (l listTodos) getTodoLists(c context.Context) (vms []viewmodels.TodoList, err error) {
	todoLists, err := l.queries.GetTodoLists(c)
	if err != nil {
		return
	}

	return viewmodels.NewTodoLists(todoLists)
}
