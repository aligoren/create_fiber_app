package commands

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

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
		return fmt.Errorf("error: %s is already contains a go.mod file", dirName)
	}

	err := createApp(name, dirName)
	if err != nil {
		return err
	}

	err = createExampleFiles(dirName)

	if err != nil {
		return err
	}

	return nil
}

var NewProjectCommand = cli.Command{
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
