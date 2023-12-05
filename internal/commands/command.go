package commands

type Command interface {
	Execute(args []string)
}

type CommandArgs struct {
	Alias string
	Args  []string
}

func NewCommandArgs(cmd string, args []string) *CommandArgs {
	return &CommandArgs{
		Alias: cmd,
		Args:  args,
	}
}
