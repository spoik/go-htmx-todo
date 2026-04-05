package deletetodo

import (
	"github.com/spoik/go-htmx-todo/internal/database/queries"
	"net/http"
)

type deleteTodo struct {
	queries *queries.Queries
}

func New(q *queries.Queries) deleteTodo {
	return deleteTodo{queries: q}
}

func (d deleteTodo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}
