package commands

import (
	"create_fiber_app/internal"
	"create_fiber_app/templates"
	"fmt"
)

// create example files with in a structure
func createExampleFiles(dirName string) error {
	filesToGenerate := []internal.ExampleFileGenerator{
		{
			Name:    fmt.Sprintf("%s%s", dirName, internal.ENTY_FILE),
			Content: templates.MAIN_TEMPLATE,
		},
		{
			Name:    fmt.Sprintf("%s%s", dirName, internal.EXAMPLE_MIDDLEWARE_FILE),
			Content: templates.EXAMPLE_MIDDLEWARE_TEMPLATE,
		},
	}

	err := internal.GenerateFiles(filesToGenerate)

	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	return nil
}
