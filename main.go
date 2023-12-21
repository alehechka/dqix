package main

import (
	"log"
	"os"

	"dqix/cmd"
)

// Version of application
var Version = "dev"

func main() {
	if err := cmd.App(Version).Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
