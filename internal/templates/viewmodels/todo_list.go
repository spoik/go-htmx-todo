package viewmodels

import (
	"errors"

	"github.com/spoik/go-htmx-todo/internal/database/queries"
	"github.com/spoik/go-htmx-todo/internal/server/routes"
)

type TodoList struct {
	queries.TodoList
	TodosUrl string
}

func NewTodoList(tl queries.TodoList) (todoList TodoList, err error) {
	todosUrl, err := routes.ListTodos.Reverse(
		"todoListId",
		tl.ID.String(),
	)

	if err != nil {
		return
	}

	todoList = TodoList{
		TodoList: tl,
		TodosUrl: todosUrl,
	}

	return
}

func NewTodoLists(todoLists []queries.TodoList) ([]TodoList, error) {
	todoListVms := make([]TodoList, len(todoLists))
	errs := []error{}

	for i, todoList := range todoLists {
		todoListVm, err := NewTodoList(todoList)

		if err != nil {
			errs = append(errs, err)
			continue
		}

		todoListVms[i] = todoListVm
	}

	if len(errs) > 0 {
		return []TodoList{}, errors.Join(errs...)
	}

	return todoListVms, nil
}
