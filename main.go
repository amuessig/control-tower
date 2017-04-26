package main

import (
	"os"

	"github.com/engineerbetter/concourse-up/commands"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "Concourse-Up"
	app.Usage = "A CLI tool to deploy Concourse CI"
	app.Commands = commands.Commands
	app.Flags = commands.GlobalFlags
	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
