package staticroute

import (
	"fmt"
	"strings"

	"github.com/spoik/go-htmx-todo/internal/server/routes/route"
)

// A static http staticRoute. Contains no substitutions. For example "GET /todos".
type staticRoute struct {
	route.Route
}

// Creates a static http route. Contains no substitutions. For example "GET /catalog".
func New(verb string, path string) staticRoute {
	if strings.ContainsAny(path, "{}") {
		panic(
			fmt.Sprintf(
				"Route contains parameters. Use a paramRoute instead: %s %s",
				verb,
				path,
			),
		)
	}

	return staticRoute{
		route.New(verb, path),
	}
}

// Returns the path.
func (r staticRoute) Reverse() string {
	return r.Path()
}
