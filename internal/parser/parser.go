package parser

type Config struct {
	Path             string
	InputFileName    string
	DatabaseFileName string
}

type Parser interface {
	Parse() error
}
