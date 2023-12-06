package command

import "kurs_scheduler/pkg/utils"

type ClearCommand struct {
}

func NewClearCommand() *ClearCommand {
	return &ClearCommand{}
}

func (c *ClearCommand) Execute(args []string) {
	utils.ClearScreen()
}
