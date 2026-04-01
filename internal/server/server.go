package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spoik/go-htmx-todo/internal/database/queries"
	"github.com/spoik/go-htmx-todo/internal/server/listtodos"
	"github.com/spoik/go-htmx-todo/internal/server/updatetodo"
	"github.com/spoik/go-htmx-todo/internal/server/middleware"
)

type Server struct {
	handler *http.Handler
	queries *queries.Queries
}

func New(q *queries.Queries) *Server {
	mux := http.NewServeMux()

	mux.Handle("GET /", listtodos.New(q))
	mux.Handle("POST /todos/{id}", updatetodo.New(q))

	wrappedMux := middleware.LogRequests(mux)

	return &Server{
		handler: &wrappedMux,
		queries: q,
	}
}

func (s *Server) Start(port int) {
	log.Printf("Starting server on :%d\n", port)
	err := http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		*s.handler,
	)

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
