package parser

type Config struct {
	Path          string
	InputFileName string
}

type Parser interface {
	Parse() error
}
