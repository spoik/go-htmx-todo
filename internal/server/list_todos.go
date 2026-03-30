package server

import "net/http"

type ListTods struct {}

func (l *ListTods) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Needs to be implemented.", http.StatusInternalServerError)
}
