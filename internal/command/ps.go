package command

import (
	"fmt"
	"kurs_scheduler/internal/process"
	"kurs_scheduler/internal/scheduler"
)

type ProcessListCommand struct {
	Scheduler    *scheduler.Scheduler
	ErrorMessage string
}

func NewProcessListCommand(scheduler *scheduler.Scheduler) *ProcessListCommand {
	return &ProcessListCommand{
		Scheduler: scheduler,
	}
}

func (cmd *ProcessListCommand) Execute(args []string) {
	cmd.Scheduler.SetMaxBurst(10)
	cmd.Scheduler.GenerateProcesses(10)
	processes := cmd.Scheduler.GetAllProcesses()
	if len(processes) < 1 {
		fmt.Println("Процессы не сгенерированы")
		return
	}

	status := func(status process.Status) string {
		switch status {
		case process.Running:
			return "Выполнение"
		case process.Waiting:
			return "Ожидание"
		case process.Readiness:
			return "Готовность"
		default:
			return ""
		}
	}

	fmt.Println("ID\tUID\tBurst\tСтатус")
	for _, proc := range processes {
		fmt.Println(proc.ID, "\t",
			proc.UID, "\t",
			proc.Burst, "\t",
			status(proc.Status),
		)
	}
}
