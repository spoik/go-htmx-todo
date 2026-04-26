package viewmodels

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spoik/go-htmx-todo/internal/database/queries"
	"github.com/spoik/go-htmx-todo/internal/server/routes"
)

type Todo struct {
	queries.Todo
	UpdateCompleteUrl string
	DeleteUrl         string
}

func NewTodo(t queries.Todo) (todo Todo, err error) {
	id := strconv.Itoa(int(t.ID))

	updateTodoCompleteUrl, err := routes.UpdateTodoComplete.Reverse("id", id)

	if err != nil {
		return
	}

	deleteUrl, err := routes.DeleteTodo.Reverse("id", id)

	if err != nil {
		return
	}

	todo = Todo{
		Todo:              t,
		UpdateCompleteUrl: updateTodoCompleteUrl,
		DeleteUrl:         deleteUrl,
	}

	return
}

func NewTodos(todos []queries.Todo) ([]Todo, error) {
	todoVms := make([]Todo, len(todos))
	errs := []error{}

	for i, todo := range todos {
		todoVm, err := NewTodo(todo)

		if err != nil {
			errs = append(errs, err)
			continue
		}

		todoVms[i] = todoVm
	}

	if len(errs) > 0 {
		return []Todo{}, errors.Join(errs...)
	}

	return todoVms, nil
}

func (t Todo) InputId() string {
	return fmt.Sprintf("todo-%d", t.ID)
}

func (t Todo) LabelId() string {
	return "label-" + t.InputId()
}

func (t Todo) IsCompleted() bool {
	return t.Complete.Bool
}
