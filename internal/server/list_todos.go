package server

import "net/http"

type ListTodos struct {}

func (l ListTodos) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Needs to be implemented.", http.StatusInternalServerError)
}
