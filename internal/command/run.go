package command

import "kurs_scheduler/internal/scheduler"

type RunCommand struct {
	Scheduler *scheduler.Scheduler
}

func NewRunCommand(scheduler *scheduler.Scheduler) *RunCommand {
	return &RunCommand{
		Scheduler: scheduler,
	}
}

func (cmd *RunCommand) Execute(args []string) {
	if len(cmd.Scheduler.Processes) < 1 {
		return
	}
	NewClearCommand().Execute(args)

	cmd.Scheduler.Run()
}
