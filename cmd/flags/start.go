package flags

import (
	"github.com/urfave/cli/v2"
)

var StartFlags = []cli.Flag{
	FlagPort,
	FlagDataPath,
}

// ArgPort is the CLI argument for the port to listen on
const ArgPort string = "port"

// FlagPort is the urfave/cli Flag configuration for the port to listen on
var FlagPort = &cli.IntFlag{
	Name:    ArgPort,
	Usage:   "Specifies the port to listen on.",
	EnvVars: []string{"PORT"},
	Value:   8080,
}

const ArgDataPath string = "data-path"

var FlagDataPath = &cli.StringFlag{
	Name:  ArgDataPath,
	Usage: "Specifies the base path to the data folder to read into memory",
	Value: "./web/data",
}
