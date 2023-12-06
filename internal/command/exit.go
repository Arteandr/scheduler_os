package command

import (
	"os"
)

type ExitCommand struct {
}

func NewExitCommand() *ExitCommand {
	return &ExitCommand{}
}

func (c *ExitCommand) Execute(args []string) {
	os.Exit(0)
}
