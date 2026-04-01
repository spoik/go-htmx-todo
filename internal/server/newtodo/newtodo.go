package newtodo

import (
	"github.com/a-h/templ"
	"github.com/spoik/go-htmx-todo/internal/templates"
)

var NewTodo = templ.Handler(templates.NewTodo())
