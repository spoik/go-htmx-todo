package dynamicroute

import (
	"fmt"
	"strings"

	"github.com/spoik/go-htmx-todo/internal/server/routes/route"
)

// A dynamic http route. Contains substitutions. For example "GET /todo/{id}".
type dynamicRoute struct {
	route.Route
}

// Creates a dynamic http route. Contains substitutions. For example "GET /todo/{id}".
func New(verb string, path string) dynamicRoute {
	if !strings.ContainsAny(path, "{}") {
		panic(
			fmt.Sprintf(
				"Route does not contain any parameters. Use a staticRoute instead: %s %s",
				verb,
				path,
			),
		)
	}

	return dynamicRoute{
		route.New(verb, path),
	}
}

// Returns the path with params replaced with values.
func (r dynamicRoute) Reverse(params ...string) (string, error) {
	replacer := r.createReverseReplacer(params...)

	path := replacer.Replace(r.Path())

	// Check if there are any params in the path that haven't been substituted.
	if strings.ContainsAny(path, "{}") {
		return "", fmt.Errorf(
			"Route reverse failed. Not all params were replaced in the following: %s",
			path,
		)
	}

	return path, nil
}

func (r dynamicRoute) createReverseReplacer(params ...string) *strings.Replacer {
	replacements := make([]string, len(params))

	for i := 0; i < len(params); i += 2 {
		paramName := params[i]
		paramValue := params[i+1]

		paramPattern := fmt.Sprintf("{%s}", paramName)
		replacements = append(replacements, paramPattern)
		replacements = append(replacements, paramValue)
	}

	return strings.NewReplacer(replacements...)
}
