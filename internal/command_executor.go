package internal

import "os/exec"

type Commands struct {
	Name          string
	Args          []string
	BeforeMessage string
	AfterMessage  string
}

type CommandData struct {
	Output []byte
	Error  error
}

func RunCommand(ch chan<- CommandData, name string, args ...string) {
	cmd := exec.Command(name, args...)
	data, err := cmd.CombinedOutput()
	ch <- CommandData{
		Error:  err,
		Output: data,
	}
}
