package cmd

import (
	"github.com/urfave/cli/v2"
)

// App represents the CLI application
func App(version string) *cli.App {
	app := cli.NewApp()
	app.Name = "Dragon Quest IX Wiki"
	app.Version = version
	app.Usage = "New wiki site for Dragon Quest IX"
	app.Commands = []*cli.Command{
		ScrapeCommand,
		ParseCommand,
	}

	return app
}
