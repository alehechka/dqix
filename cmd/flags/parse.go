package flags

import "github.com/urfave/cli/v2"

var ParseWikidotFlags = []cli.Flag{
	FlagWikidotPath,
	FlagInputDataFileName,
	FlagDatabaseFile,
}

const ArgWikidotPath string = "wikidot-path"

var FlagWikidotPath = &cli.StringFlag{
	Name:  ArgWikidotPath,
	Usage: "Specifies the path to the Wikidot data folder",
	Value: "./data/wikidot/",
}

const ArgInputDataFileName string = "input-file"

var FlagInputDataFileName = &cli.StringFlag{
	Name:  ArgInputDataFileName,
	Usage: "Name of file to use as raw input data",
	Value: "raw.json",
}
