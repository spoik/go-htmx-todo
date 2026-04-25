package routes

import "net/http"

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
