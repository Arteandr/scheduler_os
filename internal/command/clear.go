package command

import (
	"os"
	"os/exec"
)

type ClearCommand struct {
}

func NewClearCommand() *ClearCommand {
	return &ClearCommand{}
}

func (c *ClearCommand) Execute(args []string) {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
