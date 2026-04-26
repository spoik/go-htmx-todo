package updatetodocomplete

import (
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spoik/go-htmx-todo/internal/database/queries"
	"github.com/spoik/go-htmx-todo/internal/templates"
	"github.com/spoik/go-htmx-todo/internal/templates/viewmodels"
)

type updateTodoComplete struct {
	queries *queries.Queries
}

func New(q *queries.Queries) updateTodoComplete {
	return updateTodoComplete{queries: q}
}

func (u updateTodoComplete) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	origTodo, err := u.getTodo(w, r)

	if err != nil {
		templates.GenericError(w, r, err)
		return
	}

	updatedTodo, err := u.updateTodoComplete(w, r, origTodo)

	if err != nil {
		u.genericHTMLError(w, r, origTodo, err)
		return
	}

	templates.Render(w, r, templates.Todo(updatedTodo, ""))
}

func (u updateTodoComplete) getTodo(w http.ResponseWriter, r *http.Request) (todo queries.Todo, err error) {
	id := r.PathValue("id")

	idInt, err := strconv.ParseInt(id, 10, 32)

	if err != nil {
		return
	}

	return u.queries.GetTodo(r.Context(), int32(idInt))
}

func (u updateTodoComplete) updateTodoComplete(w http.ResponseWriter, r *http.Request, todo queries.Todo) (todoVm viewmodels.Todo, err error) {
	var complete pgtype.Bool
	if err = complete.Scan(r.FormValue("complete") == "on"); err != nil {
		return
	}

	params := queries.UpdateTodoCompleteParams{
		ID:       todo.ID,
		Complete: complete,
	}

	updatedTodo, err := u.queries.UpdateTodoComplete(r.Context(), params)
	if err != nil {
		return
	}

	return viewmodels.NewTodo(updatedTodo)
}

func (u updateTodoComplete) genericHTMLError(w http.ResponseWriter, r *http.Request, todo queries.Todo, err error) {
	todoVm, err := viewmodels.NewTodo(todo)

	if err != nil {
		templates.GenericError(w, r, err)
		return
	}

	templates.Render(w, r, templates.Todo(todoVm, templates.GenericErrorMessage))
}
