package parser

type Config struct {
	Path            string
	InputFileName   string
	OutputDirectory string
}

type Parser interface {
	Parse() error
}
