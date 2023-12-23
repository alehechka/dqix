package cmd

import (
	"dqix/cmd/flags"
	"dqix/internal/router"

	"github.com/urfave/cli/v2"
)

func start(ctx *cli.Context) (err error) {
	return router.NewRouter().Run(ctx.Int(flags.PortArg))
}

// StartCommand starts the application.
var StartCommand = &cli.Command{
	Name:   "start",
	Usage:  "Start the application.",
	Action: start,
	Flags:  flags.StartFlags,
}
