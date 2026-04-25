package cancelnewtodolist

import (
	"net/http"

	"github.com/spoik/go-htmx-todo/internal/templates"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	templates.Render(w, r, templates.AddTodoList())
}
