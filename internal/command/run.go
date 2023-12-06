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
	NewClearCommand().Execute(args)

	cmd.Scheduler.Run()
}
