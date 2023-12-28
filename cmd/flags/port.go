package flags

import (
	"github.com/urfave/cli/v2"
)

// PortArg is the CLI argument for the port to listen on
const PortArg string = "port"

// PortFlag is the urfave/cli Flag configuration for the port to listen on
var PortFlag = &cli.IntFlag{
	Name:    PortArg,
	Usage:   "Specifies the port to listen on.",
	EnvVars: []string{"PORT"},
	Value:   8080,
}
