package main

import (
	"create_fiber_app/commands"
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {

	// These settings can be changed
	app := &cli.App{
		Name:        "Create Fiber App",
		Usage:       "Helps your Fiber development easy",
		Description: "Creates Fiber apps with no build configuration",
		Version:     "1.0",
		Commands: []*cli.Command{
			&commands.NewProjectCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
