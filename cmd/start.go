package cmd

import (
	"dqix/cmd/flags"
	"dqix/internal/router"

	"github.com/urfave/cli/v2"
)

func start(ctx *cli.Context) (err error) {
	return router.NewRouter(router.WithData(ctx.String(flags.ArgDataPath))).Run(ctx.Int(flags.ArgPort))
}

// StartCommand starts the application.
var StartCommand = &cli.Command{
	Name:   "start",
	Usage:  "Start the application.",
	Action: start,
	Flags:  flags.StartFlags,
}
