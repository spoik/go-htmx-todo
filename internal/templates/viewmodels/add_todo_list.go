package viewmodels

import "github.com/spoik/go-htmx-todo/internal/server/routes"

type AddTodoList struct {
	Url string
}

func NewAddTodoList() (AddTodoList, error) {
	url, err := routes.NewTodoList.Reverse()

	if err != nil {
		return AddTodoList{}, err
	}

	return AddTodoList{Url: url}, nil
}
