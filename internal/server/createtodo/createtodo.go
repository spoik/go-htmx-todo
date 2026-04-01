package createtodo

import (
	"net/http"

	"github.com/spoik/go-htmx-todo/internal/templates"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	// TODO: Validate new todo
	// TODO: Try to write the new todo to the database
	templates.TodoForm("The Title", "The Error").Render(r.Context(), w)
}
