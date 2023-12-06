package command

import (
	"fmt"
	"kurs_scheduler/internal/process"
	"kurs_scheduler/internal/scheduler"
	"kurs_scheduler/pkg/utils"
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
	processes := cmd.Scheduler.GetAllProcesses()
	if len(processes) < 1 {
		utils.Error("Процессы не сгенерированы")
		return
	}

	status := func(status process.Status) string {
		switch status {
		case process.Running:
			return "Выполнение"
		case process.Completed:
			return "Выполнен"
		case process.Waiting:
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
