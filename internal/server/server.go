package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spoik/go-htmx-todo/internal/database/queries"
	"github.com/spoik/go-htmx-todo/internal/server/createtodo"
	"github.com/spoik/go-htmx-todo/internal/server/deletetodo"
	"github.com/spoik/go-htmx-todo/internal/server/listtodos"
	"github.com/spoik/go-htmx-todo/internal/server/middleware"
	"github.com/spoik/go-htmx-todo/internal/server/newtodo"
	"github.com/spoik/go-htmx-todo/internal/server/routes"
	"github.com/spoik/go-htmx-todo/internal/server/updatetodocomplete"
)

type Server struct {
	handler *http.Handler
}

func New(q *queries.Queries) *Server {
	mux := http.NewServeMux()

	mux.Handle(routes.ListTodos.Pattern(), listtodos.New(q))
	mux.Handle(routes.NewTodo.Pattern(), newtodo.NewTodo)
	mux.Handle(routes.UpdateTodoComplete.Pattern(), updatetodocomplete.New(q))
	mux.Handle(routes.CreateTodo.Pattern(), createtodo.New(q))
	mux.Handle(routes.DeleteTodo.Pattern(), deletetodo.New(q))

	wrappedMux := middleware.LogRequests(mux)

	return &Server{handler: &wrappedMux}
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
