package routes

import (
	"net/http"

	"github.com/spoik/go-htmx-todo/internal/server/routes/staticroute"
)

var NewTodoList = staticroute.New(http.MethodGet, "/lists/new")

var CancelNewTodoList = staticroute.New(http.MethodGet, "/lists/new/cancel")

var CreateTodoList = staticroute.New(http.MethodPost, "/lists")
