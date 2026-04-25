package routes

import (
	"net/http"
	"github.com/spoik/go-htmx-todo/internal/server/routes/staticroute"
)

var Root = staticroute.New(http.MethodGet, "/")
