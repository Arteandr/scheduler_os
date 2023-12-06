package command

import (
	"fmt"
	"kurs_scheduler/internal/scheduler"
)

type RunCommand struct {
	Scheduler *scheduler.Scheduler
}

func NewRunCommand(scheduler *scheduler.Scheduler) *RunCommand {
	return &RunCommand{
		Scheduler: scheduler,
	}
}

func (cmd *RunCommand) Execute(args []string) {
	if cmd.Scheduler.Quantum < 1 {
		fmt.Println("Не указан квант выполнения")
		return
	} else if cmd.Scheduler.MaxBurst < 1 {
		fmt.Println("Не указан максимальный CPU Burst")
		return
	}
	NewClearCommand().Execute(args)

	cmd.Scheduler.Run()
	cmd.Scheduler.RunDraw()
}
