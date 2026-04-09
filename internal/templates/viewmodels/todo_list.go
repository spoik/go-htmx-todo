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

func NewTodoList(tl queries.TodoList) (TodoList, error) {
	todosUrl, err := routes.ListTodos.Reverse("listID", tl.ID.String())
	if err != nil {
		return TodoList{}, err
	}
	return TodoList{
		TodoList: tl,
		TodosUrl: todosUrl,
	}, nil
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
