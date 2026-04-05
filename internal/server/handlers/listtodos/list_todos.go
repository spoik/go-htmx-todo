package listtodos

import (
	"net/http"

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
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	todos, err := l.queries.GetTodos(r.Context())

	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	todoVms, err := viewmodels.NewTodos(todos)

	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	templates.Todos(todoVms).Render(r.Context(), w)
}
