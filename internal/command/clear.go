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
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
