package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spoik/go-htmx-todo/internal/database"
	"github.com/spoik/go-htmx-todo/internal/templates"
)

func Create() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /", ListTods{})
	mux.Handle("POST /todos/{id}", UpdateTodo{})

	return mux
}

func Start(mux *http.ServeMux, port int) {
	log.Printf("Starting server on :%d\n", port)
	err := http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		mux,
	)

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
