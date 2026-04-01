package viewmodels

import (
	"fmt"
	"strconv"

	"github.com/spoik/go-htmx-todo/internal/database/queries"
	"github.com/spoik/go-htmx-todo/internal/server/routes"
)

type Todo struct {
	todo   queries.Todo
	putUrl string
}

// TODO: Return an optional error to handle when route reversal fails
func NewTodo(t queries.Todo) Todo {
	putUrl := putURL(t)

	return Todo{
		todo:   t,
		putUrl: putUrl,
	}
}

func putURL(t queries.Todo) string {
	idStr := strconv.Itoa(int(t.ID))
	return routes.ToggleTodoComplete.Reverse("id", idStr)
}

func (t Todo) InputId() string {
	return fmt.Sprintf("todo-%d", t.todo.ID)
}

func (t Todo) LabelId() string {
	return "label-" + t.InputId()
}

func (t Todo) IsChecked() bool {
	return t.todo.Complete.Bool
}

func (t Todo) TodoTitle() string {
	return t.todo.Title
}
func (t Todo) PutURL() string {
	return t.putUrl
}
