package cmd

import (
	"dqix/cmd/flags"
	"dqix/internal/parser"
	"dqix/internal/parser/wikidot"

	"github.com/urfave/cli/v2"
)

func parseWikidot(ctx *cli.Context) (err error) {
	return wikidot.Init(&parser.Config{
		Path:            ctx.String(flags.ArgWikidotPath),
		InputFileName:   ctx.String(flags.ArgInputDataFileName),
		OutputDirectory: ctx.String(flags.ArgOutputDirectory),
	}).Parse()
}

var parseWikidotCommand = &cli.Command{
	Name:   "wikidot",
	Usage:  "Parse Wikidot data into standard format",
	Action: parseWikidot,
	Flags:  flags.ParseWikidotFlags,
}

var ParseCommand = &cli.Command{
	Name:  "parse",
	Usage: "Parse a specific data format into standard",
	Subcommands: []*cli.Command{
		parseWikidotCommand,
	},
}
