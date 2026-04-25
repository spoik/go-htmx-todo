package routes

import "net/http"

var Root = route{
	verb: http.MethodGet,
	path: "/",
}

