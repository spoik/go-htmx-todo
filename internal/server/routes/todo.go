package routes

import (
	"net/http"

	"github.com/spoik/go-htmx-todo/internal/server/routes/dynamicroute"
	"github.com/spoik/go-htmx-todo/internal/server/routes/staticroute"
)

var ListTodos = dynamicroute.New(http.MethodGet, "/lists/{todoListId}")

var UpdateTodoComplete = dynamicroute.New(http.MethodPut, "/todo/{id}/complete")

var NewTodo = dynamicroute.New(http.MethodGet, "/lists/{todoListId}/todo/new")

var CreateTodo = staticroute.New(http.MethodPost, "/todo")

var DeleteTodo = dynamicroute.New(http.MethodDelete, "/todo/{id}")
