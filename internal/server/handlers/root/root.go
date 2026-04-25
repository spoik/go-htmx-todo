package root

import (
	"net/http"

	"github.com/spoik/go-htmx-todo/internal/database/queries"
	"github.com/spoik/go-htmx-todo/internal/server/response"
	"github.com/spoik/go-htmx-todo/internal/server/routes"
)

type root struct {
	queries *queries.Queries
}

func New(q *queries.Queries) root {
	return root{queries: q}
}

func (h root) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	todoLists, err := h.queries.GetTodoLists(r.Context())
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	if len(todoLists) == 0 {
		http.NotFound(w, r)
		return
	}

	firstList := todoLists[0]
	url, err := routes.ListTodos.Reverse("todoListId", firstList.ID.String())
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	http.Redirect(w, r, url, http.StatusSeeOther)
}
