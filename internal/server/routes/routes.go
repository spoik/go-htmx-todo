package routes

var ListTodos = route{
	verb: "GET",
	path: "/",
}

var ToggleTodoComplete = route{
	verb: "PUT",
	path: "/todo/{id}/togglecomplete",
}
