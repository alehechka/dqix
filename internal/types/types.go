package types

type PageContent struct {
	Path  string
	Text  []string
	Links map[string]string
}

type Pages map[string]PageContent
