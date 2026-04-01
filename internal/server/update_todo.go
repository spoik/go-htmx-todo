package server

import (
	// "log"
	"net/http"
	"strconv"

	"github.com/spoik/go-htmx-todo/internal/database/queries"
	// "github.com/spoik/go-htmx-todo/internal/server/response"
	// "github.com/spoik/go-htmx-todo/internal/templates"
)

type UpdateTodo struct {
	queries *queries.Queries
}

func (u UpdateTodo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	_, err := strconv.ParseInt(id, 10, 32)

	if err != nil {
		http.Error(
			w,
			"Invalid todo id. Must be an integer.",
			http.StatusUnprocessableEntity,
		)
		return
	}

	// todo, err := u.queries.GetTodo(r.Context(), int32(idInt))
	// log.Printf("%v", err)
	//
	// if err != nil {
	// 	response.InternalServerError(w, r, err)
	// 	return
	// }
	//
	// todo.Complete.Bool = !todo.Complete.Bool
	//
	// templates.Todo(todo).Render(r.Context(), w)
}
