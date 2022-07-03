package internal

import (
	"os"
	"path/filepath"
)

func createFile(path string, content string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0770); err != nil {
		return nil, err
	}

	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	file.WriteString(content)

	return file, nil
}

type ExampleFileGenerator struct {
	Name    string
	Content string
}

func (f ExampleFileGenerator) GenerateFile() (*os.File, error) {
	return createFile(f.Name, f.Content)
}

func GenerateFiles(files []ExampleFileGenerator) error {
	for _, file := range files {
		_, err := file.GenerateFile()
		if err != nil {
			return err
		}
	}

	return nil
}
