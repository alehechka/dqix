package fn

import (
	"path"
	"strings"

	"github.com/a-h/templ"
)

func Path(parts ...string) templ.SafeURL {
	path := path.Join(parts...)

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	return templ.URL(path)
}
