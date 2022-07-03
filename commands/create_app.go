package commands

import (
	"create_fiber_app/internal"
	"fmt"
	"os"
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
