package newtodo

import (
	"net/http"

	"github.com/spoik/go-htmx-todo/internal/templates"
)

func NewTodo(w http.ResponseWriter, r *http.Request) {
	todoListId := r.PathValue("todoListId")
	if todoListId == "" {
		http.NotFound(w, r)
		return
	}
	templates.Render(w, r, templates.TodoForm(todoListId, "", ""))
}
