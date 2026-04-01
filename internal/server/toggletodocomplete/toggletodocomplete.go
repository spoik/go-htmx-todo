package toggletodocomplete

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spoik/go-htmx-todo/internal/database/queries"
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

	todoVm, err := viewmodels.NewTodo(todo)

	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	// TODO: Handle errors returned from rendering. Test if error handling can be done in the template and then returned here.
	templates.Todo(todoVm).Render(r.Context(), w)
}

func (u toggleTodoComplete) getTodo(w http.ResponseWriter, r *http.Request) (queries.Todo, error) {
	id := r.PathValue("id")

	idInt, err := strconv.ParseInt(id, 10, 32)

	if err != nil {
		http.Error(
			w,
			"Invalid todo id. Must be an integer.",
			http.StatusUnprocessableEntity,
		)
		return queries.Todo{}, err
	}

	todo, err := u.queries.GetTodo(r.Context(), int32(idInt))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.NotFound(w, r)
		} else {
			response.InternalServerError(w, r, err)
		}

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

	todo, err := u.queries.UpdateTodoComplete(r.Context(), params)

	if err != nil {
		response.InternalServerError(w, r, err)
		return queries.Todo{}, err
	}

	return todo, nil
}
