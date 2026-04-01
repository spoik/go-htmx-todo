package routes

import (
	"fmt"
	"strings"
)

type route struct {
	verb string
	path string
}

// Returns the route with the verb. Designed to be passed to a server mux.
func (r route) Pattern() string {
	return fmt.Sprintf("%s %s", r.verb, r.path)
}

// Returns the path with params replaced with values.
func (r route) Reverse(params ...string) (string, error) {
	replacer := r.createReverseReplacer(params...)

	path := replacer.Replace(r.path)

	// Check if there are any params in the path that haven't been substituted.
	if strings.ContainsAny(path, "{}") {
		return "", fmt.Errorf(
			"Route reverse failed. Not all params were replaced in the following: %s",
			path,
		)
	}

	return path, nil
}

func (r route) createReverseReplacer(params ...string) *strings.Replacer {
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
