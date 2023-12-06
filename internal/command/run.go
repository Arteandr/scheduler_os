package command

import (
	"kurs_scheduler/internal/scheduler"
	"kurs_scheduler/pkg/utils"
	"strconv"
)

type RunCommand struct {
	Scheduler    *scheduler.Scheduler
	ErrorMessage string
}

func NewRunCommand(scheduler *scheduler.Scheduler) *RunCommand {
	return &RunCommand{
		Scheduler:    scheduler,
		ErrorMessage: "Использование run <count>",
	}
}

func (cmd *RunCommand) Execute(args []string) {
	switch {
	case cmd.Scheduler.Quantum < 1:
		utils.Error("Не указан квант исполнения")
		return
	case cmd.Scheduler.MaxBurst < 1:
		utils.Error("Не указан максимальный CPU Burst")
		return
	case len(args) != 1:
		utils.Error(cmd.ErrorMessage)
		return
	}
	NewClearCommand().Execute(args)

	var count int
	var err error
	count, err = strconv.Atoi(args[0])
	if err != nil || count < 1 || count > 15 {
		utils.Error(cmd.ErrorMessage)
		return
	}

	cmd.Scheduler.Run(count)
	cmd.Scheduler.RunDraw()
}
