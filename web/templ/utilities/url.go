package fn

import (
	"strings"

	"github.com/a-h/templ"
)

func Path(parts ...string) templ.SafeURL {
	path := strings.Join(parts, "/")

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	return templ.URL(path)
}
