package createtodo

import "net/http"

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todo created"))
}
