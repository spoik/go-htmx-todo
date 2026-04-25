package route

import "fmt"

type Route struct {
	verb string
	path string
}

func New(verb string, path string) Route {
	// TODO: Check that the verb is a valid http verb. If it is, return an error.

	// TODO: Check that the path is a valid path. If it is, return an error
	return Route{
		verb: verb,
		path: path,
	}

}

// Returns the route with the verb. Designed to be passed to a server mux.
func (r Route) Pattern() string {
	return fmt.Sprintf("%s %s", r.verb, r.Path())
}

func (r Route) Path() string {
	return r.path
}
