package cmd

import (
	"database/sql"
	"dqix/cmd/flags"
	ddl "dqix/db"
	"dqix/internal/parser"
	"dqix/internal/parser/wikidot"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"

	"github.com/urfave/cli/v2"
)

func parseWikidot(ctx *cli.Context) (err error) {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		return err
	}

	// create tables
	if _, err := db.ExecContext(ctx.Context, ddl.DDL); err != nil {
		return err
	}

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
