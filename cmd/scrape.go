package cmd

import (
	"dqix/cmd/flags"
	"fmt"

	"github.com/urfave/cli/v2"
)

func scrapeWikidot(ctx *cli.Context) (err error) {
	fmt.Println("scraping...")
	return
}

var scrapeWikidotCommand = &cli.Command{
	Name:   "wikidot",
	Usage:  "Scrape Wikidot for data",
	Action: scrapeWikidot,
	Flags:  flags.ScrapeWikidotFlags,
}

var ScrapeCommand = &cli.Command{
	Name:  "scrape",
	Usage: "Scrape a specific webpage for data",
	Subcommands: []*cli.Command{
		scrapeWikidotCommand,
	},
}
