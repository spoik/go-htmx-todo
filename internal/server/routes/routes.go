package routes

import "net/http"

var ListTodos = route{
	verb: http.MethodGet,
	path: "/",
}

var ToggleTodoComplete = route{
	verb: http.MethodPut,
	path: "/todo/{id}/togglecomplete",
}

var NewTodo = route{
	verb: http.MethodGet,
	path: "/todo/new",
}

var CreateTodo = route{
	verb: http.MethodPost,
	path: "/todo",
}
