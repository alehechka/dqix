package wikidot

import (
	"dqix/internal/parser"
)

type WikidotParser struct {
	config *parser.Config
}

func Init(config *parser.Config) parser.Parser {
	return &WikidotParser{
		config: config,
	}
}

func (p WikidotParser) Parse() (err error) {
	return
}
