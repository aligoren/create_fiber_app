package main

import (
	"create_fiber_app/commands"
	cli "github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:        "Create Fiber App",
		Usage:       "Helps your Fiber development easy",
		Description: "Creates Fiber apps with no build configuration",
		Version:     "1.0",
		Commands: []*cli.Command{
			&commands.CreateCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
