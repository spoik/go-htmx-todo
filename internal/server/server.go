package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/spoik/go-htmx-todo/internal/templates"
)

func create() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/", templ.Handler(templates.Todos()))

	return mux
}

func Start(port int) {
    mux := create()

	log.Printf("Starting server on :%d\n", port)
	err := http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		mux,
	)

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
