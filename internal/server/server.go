package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spoik/go-htmx-todo/internal/database/queries"
	// "github.com/spoik/go-htmx-todo/internal/templates"
)

type Server struct {
	mux     *http.ServeMux
	queries *queries.Queries
}

func New(q *queries.Queries) *Server {
	mux := http.NewServeMux()

	mux.Handle("GET /", ListTodos{})
	mux.Handle("POST /todos/{id}", UpdateTodo{})

	return &Server{
		mux:     mux,
		queries: q,
	}
}

func (s *Server) Start(port int) {
	log.Printf("Starting server on :%d\n", port)
	err := http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		s.mux,
	)

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
