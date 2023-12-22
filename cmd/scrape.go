package cmd

import (
	"dqix/cmd/flags"
	"dqix/internal/scraper"
	"dqix/internal/scraper/wikidot"
	"net/url"

	"github.com/urfave/cli/v2"
)

func scrapeWikidot(ctx *cli.Context) (err error) {
	wikiURL, err := url.Parse(ctx.String(flags.ArgWikidotURL))
	if err != nil {
		return err
	}

	return wikidot.Init(&scraper.Config{
		WikiURL: wikiURL,
	}).Scrape()

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
