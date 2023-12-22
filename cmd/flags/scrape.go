package flags

import "github.com/urfave/cli/v2"

var ScrapeWikidotFlags = []cli.Flag{
	FlagWikidotURL,
}

const ArgWikidotURL string = "wikidot-url"

var FlagWikidotURL = &cli.StringFlag{
	Name:  ArgWikidotURL,
	Usage: "Specifies the URL to use when requesting against Wikidot",
	Value: "http://dqnine.wikidot.com",
}
