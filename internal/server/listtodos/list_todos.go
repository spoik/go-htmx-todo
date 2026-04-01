package listtodos

import (
	"net/http"

	"github.com/spoik/go-htmx-todo/internal/database/queries"
	"github.com/spoik/go-htmx-todo/internal/server/response"
	"github.com/spoik/go-htmx-todo/internal/templates"
)

type ListTodos struct {
	queries *queries.Queries
}

func New(q *queries.Queries) ListTodos {
	return ListTodos{queries: q}
}

func (l ListTodos) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	todos, err := l.queries.GetTodos(r.Context())

	if err != nil {
		response.InternalServerError(w, err)
		return
	}

	templates.Todos(todos).Render(r.Context(), w)
}
