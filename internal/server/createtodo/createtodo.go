package createtodo

import "net/http"

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	// TODO: Validate new todo
	// TODO: Try to write the new todo to the database
	w.Write([]byte("todo created"))
}
