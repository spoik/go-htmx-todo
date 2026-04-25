package routes

import "net/http"

var Root = route{
	verb: http.MethodGet,
	path: "/",
}

var ListTodos = route{
	verb: http.MethodGet,
	path: "/lists/{todoListId}",
}

var UpdateTodoComplete = route{
	verb: http.MethodPut,
	path: "/todo/{id}/complete",
}

var NewTodo = route{
	verb: http.MethodGet,
	path: "/lists/{todoListId}/todo/new",
}

var CreateTodo = route{
	verb: http.MethodPost,
	path: "/todo",
}

var DeleteTodo = route{
	verb: http.MethodDelete,
	path: "/todo/{id}",
}

var NewTodoList = route{
	verb: http.MethodGet,
	path: "/lists/new",
}

var CancelNewTodoList = route{
	verb: http.MethodGet,
	path: "/lists/new/cancel",
}

var CreateTodoList = route{
	verb: http.MethodPost,
	path: "/lists",
}
