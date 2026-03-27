package server

import (
	"net/http"

	"github.com/spoik/go-htmx-todo/internal/db"
	"github.com/spoik/go-htmx-todo/internal/templates"
)

func UpdateTodo(todos *[]db.Todo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		if todos == nil {
			http.Error(w, "Todos \"database\" is nil.", http.StatusInternalServerError)
			return
		}

		var todo *db.Todo

		for i, t := range *todos {
			if t.ID == id {
				todo = &(*todos)[i]
			}
		}

		if todo == nil {
			http.NotFound(w, r)
			return
		}

		todo.Complete = !todo.Complete

		templates.Todo(*todo).Render(r.Context(), w)
		return
	}
}
