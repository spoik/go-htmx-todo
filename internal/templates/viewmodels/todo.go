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
	ToggleCompleteUrl string
}

func NewTodo(t queries.Todo) (Todo, error) {
	putUrl, err := toggleCompleteUrl(t)

	if err != nil {
		return Todo{}, err
	}

	return Todo{
		Todo:              t,
		ToggleCompleteUrl: putUrl,
	}, nil
}

func toggleCompleteUrl(t queries.Todo) (string, error) {
	idStr := strconv.Itoa(int(t.ID))
	return routes.ToggleTodoComplete.Reverse("id", idStr)
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

func (t Todo) IsChecked() bool {
	return t.Complete.Bool
}
