package listtodos

import (
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
	todoListIdStr := r.PathValue("todoListId")
	if todoListIdStr == "" {
		http.NotFound(w, r)
		return
	}

	var todoListId pgtype.UUID
	err := todoListId.Scan(todoListIdStr)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	todos, err := l.getTodos(r, todoListId)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	todoLists, err := l.getTodoLists(r)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	templates.SidebarAndTodos(todoListIdStr, todos, todoLists).Render(r.Context(), w)
}

func (l listTodos) getTodos(r *http.Request, todoListId pgtype.UUID) ([]viewmodels.Todo, error) {
	todos, err := l.queries.GetTodos(r.Context(), todoListId)
	if err != nil {
		return []viewmodels.Todo{}, err
	}

	todoVms, err := viewmodels.NewTodos(todos)
	if err != nil {
		return []viewmodels.Todo{}, err
	}
	return todoVms, nil
}

func (l listTodos) getTodoLists(r *http.Request) ([]viewmodels.TodoList, error) {
	todoLists, err := l.queries.GetTodoLists(r.Context())
	if err != nil {
		return []viewmodels.TodoList{}, err
	}

	todoListVms, err := viewmodels.NewTodoLists(todoLists)
	if err != nil {
		return []viewmodels.TodoList{}, err
	}
	return todoListVms, nil
}
