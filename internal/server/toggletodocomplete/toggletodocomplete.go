package toggletodocomplete

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spoik/go-htmx-todo/internal/database/queries"
	"github.com/spoik/go-htmx-todo/internal/server/log"
	"github.com/spoik/go-htmx-todo/internal/server/response"
	"github.com/spoik/go-htmx-todo/internal/templates"
	"github.com/spoik/go-htmx-todo/internal/templates/viewmodels"
)

type toggleTodoComplete struct {
	queries *queries.Queries
}

func New(q *queries.Queries) toggleTodoComplete {
	return toggleTodoComplete{queries: q}
}

func (u toggleTodoComplete) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	todo, err := u.getTodo(w, r)

	if err != nil {
		return
	}

	todo, err = u.toggleTodoComplete(w, r, todo)

	if err != nil {
		return
	}

	u.renderUpdatedTodo(w, r, todo)
}

func (u toggleTodoComplete) getTodo(w http.ResponseWriter, r *http.Request) (queries.Todo, error) {
	id := r.PathValue("id")

	idInt, err := strconv.ParseInt(id, 10, 32)

	if err != nil {
		error := errors.New("Invalid todo id. Must be an integer.")
		templates.GenericHTMLError(w, r, error)
		return queries.Todo{}, err
	}

	todo, err := u.queries.GetTodo(r.Context(), int32(idInt))

	if err != nil {
		templates.GenericHTMLError(w, r, err)

		return queries.Todo{}, err
	}

	return todo, nil
}

func (u toggleTodoComplete) toggleTodoComplete(w http.ResponseWriter, r *http.Request, todo queries.Todo) (queries.Todo, error) {
	params := queries.UpdateTodoCompleteParams{
		ID: todo.ID,
		Complete: pgtype.Bool{
			Bool:  !todo.Complete.Bool,
			Valid: true,
		},
	}

	updatedTodo, err := u.queries.UpdateTodoComplete(r.Context(), params)

	if err != nil {
		u.genericHTMLError(w, r, todo, err)
		return queries.Todo{}, err
	}

	return updatedTodo, nil
}

func (u toggleTodoComplete) renderUpdatedTodo(w http.ResponseWriter, r *http.Request, todo queries.Todo) {
	todoVm, err := viewmodels.NewTodo(todo)

	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	templates.RenderOrInternalError(w, r, templates.Todo(todoVm, ""))
}

func (u toggleTodoComplete) genericHTMLError(w http.ResponseWriter, r *http.Request, todo queries.Todo, err error) {
	log.UnhandledError(r, err)
	todoVm, err := viewmodels.NewTodo(todo)

	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	templates.RenderOrInternalError(w, r, templates.Todo(todoVm, templates.GenericErrorMessage))
}
