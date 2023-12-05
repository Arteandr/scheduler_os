package command

type RunCommand struct {
}

func NewRunCommand() *RunCommand {
	return &RunCommand{}
}

func (c *RunCommand) Execute(args []string) {
}
