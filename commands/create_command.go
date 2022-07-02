package commands

import (
	"create_fiber_app/internal"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
)

// prepare shell commands and run them sequentially
func createApp(name string, dirName string) error {

	err := os.Chdir(dirName)
	if err != nil {
		return err
	}

	commands := []internal.Commands{
		{
			Name:          "go",
			Args:          []string{"mod", "init", name},
			BeforeMessage: fmt.Sprintf("Project %s will be started", name),
			AfterMessage:  fmt.Sprintf("Project %s is created", name),
		},
		{
			Name:          "go",
			Args:          []string{"get", internal.FIBER_CURRENT_REPO},
			BeforeMessage: "Fiber is downloading...",
			AfterMessage:  "Fiber is downloaded!",
		},
	}

	for _, command := range commands {
		fmt.Printf("%s\n", command.BeforeMessage)
		commandData := make(chan internal.CommandData)

		go internal.RunCommand(commandData, command.Name, command.Args...)
		commandResponse := <-commandData

		fmt.Printf("%v\n", string(commandResponse.Output))
		fmt.Printf("%s\n", command.AfterMessage)
	}

	return nil
}

func action(context *cli.Context) error {

	name := context.String("name")
	dirName := context.String("dir")

	if len(name) == 0 {
		return errors.New("project name must be specified")
	}

	if len(dirName) == 0 {
		currentWorkingDirectory, err := os.Getwd()
		if err != nil {
			return err
		}

		dirName = currentWorkingDirectory
	}

	goModFilePath := filepath.FromSlash(fmt.Sprintf("%s/go.mod", dirName))

	fmt.Printf("%s\n", goModFilePath)

	if _, err := os.Stat(goModFilePath); err == nil {
		return errors.New(fmt.Sprintf("error: %s is already contains a go.mod file", dirName))
	}

	err := createApp(name, dirName)
	if err != nil {
		return err
	}

	return nil
}

var CreateCommand = cli.Command{
	Name:    "create",
	Aliases: []string{"c"},
	Usage:   "Creates a new project",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "Name of the your project",
			Required: true,
		}, &cli.StringFlag{
			Name:    "dir",
			Aliases: []string{"d"},
			Usage:   "Installation directory of your app",
		},
	},
	Action: action,
}
