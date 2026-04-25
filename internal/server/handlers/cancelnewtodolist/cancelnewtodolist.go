package cancelnewtodolist

import (
	"net/http"

	"github.com/spoik/go-htmx-todo/internal/server/response"
	"github.com/spoik/go-htmx-todo/internal/templates"
	"github.com/spoik/go-htmx-todo/internal/templates/viewmodels"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	addTodoList, err := viewmodels.NewAddTodoList()
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}
	templates.Render(w, r, templates.AddTodoList(addTodoList))
}
