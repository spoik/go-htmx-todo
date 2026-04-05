package routes

import "net/http"

var ListTodos = route{
	verb: http.MethodGet,
	path: "/",
}

var UpdateTodoComplete = route{
	verb: http.MethodPut,
	path: "/todo/{id}/complete",
}

var NewTodo = route{
	verb: http.MethodGet,
	path: "/todo/new",
}

var CreateTodo = route{
	verb: http.MethodPost,
	path: "/todo",
}

var DeleteTodo = route{
	verb: http.MethodDelete,
	path: "/todo/{id}",
}
