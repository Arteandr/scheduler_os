package commands

import "fmt"

type RunCommand struct {
}

func NewRunCommand() *RunCommand {
	return &RunCommand{}
}

func (r *RunCommand) Execute(args []string) {
	fmt.Println("RUN COMMAND")
}
