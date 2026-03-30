package server

import (
	"net/http"
	// "github.com/spoik/go-htmx-todo/internal/templates"
)

type UpdateTodo struct{}

func (u UpdateTodo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Needs to be implemented.", http.StatusInternalServerError)
	// id := r.PathValue("id")
	//
	// if todos == nil {
	// 	http.Error(w, "Todos \"database\" is nil.", http.StatusInternalServerError)
	// 	return
	// }
	//
	// var todo *db.Todo
	//
	// for i, t := range *todos {
	// 	if t.ID == id {
	// 		todo = &(*todos)[i]
	// 	}
	// }
	//
	// if todo == nil {
	// 	http.NotFound(w, r)
	// 	return
	// }
	//
	// todo.Complete = !todo.Complete
	//
	// templates.Todo(*todo).Render(r.Context(), w)
	// return
}
