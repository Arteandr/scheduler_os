package commands

import "fmt"

type RunCommand struct {
}

func (r *RunCommand) Execute() {
	fmt.Println("RUN COMMAND")
}
