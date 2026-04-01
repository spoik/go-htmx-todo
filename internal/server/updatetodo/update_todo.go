package updatetodo

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spoik/go-htmx-todo/internal/database/queries"
	"github.com/spoik/go-htmx-todo/internal/server/response"
	"github.com/spoik/go-htmx-todo/internal/templates"
)

type updateTodo struct {
	queries *queries.Queries
}

func New(q *queries.Queries) updateTodo {
	return updateTodo{queries: q}
}

func (u updateTodo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	todo, err := u.getTodo(w, r)

	if err != nil {
		return
	}

	todo, err = u.toggleTodoComplete(w, r, todo)

	if err != nil {
		return
	}

	templates.Todo(todo).Render(r.Context(), w)
}

func (u updateTodo) getTodo(w http.ResponseWriter, r *http.Request) (queries.Todo, error) {
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

func (u updateTodo) toggleTodoComplete(w http.ResponseWriter, r *http.Request, todo queries.Todo) (queries.Todo, error) {
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
