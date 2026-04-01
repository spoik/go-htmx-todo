package listtodos

import (
	"net/http"

	"github.com/spoik/go-htmx-todo/internal/database/queries"
	"github.com/spoik/go-htmx-todo/internal/server/response"
	"github.com/spoik/go-htmx-todo/internal/templates"
)

type listTodos struct {
	queries *queries.Queries
}

func New(q *queries.Queries) listTodos {
	return listTodos{queries: q}
}

func (l listTodos) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	todos, err := l.queries.GetTodos(r.Context())

	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	templates.Todos(todos).Render(r.Context(), w)
}
