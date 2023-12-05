package commands

type CommandExecutor struct {
	CMDs map[string]Command
}

func NewCommandExecutor() *CommandExecutor {
	cmdExecutor := &CommandExecutor{
		CMDs: make(map[string]Command),
	}

	cmdExecutor.CMDs["run"] = NewRunCommand()

	return cmdExecutor
}

func (e *CommandExecutor) ExecuteCMD(input *CommandArgs) Status {
	cmd, ok := e.CMDs[input.Alias]
	if !ok {
		return ErrorCmd
	}

	cmd.Execute(input.Args)

	return SuccessCmd
}
