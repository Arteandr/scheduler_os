package command

import (
	"fmt"
	"kurs_scheduler/internal/scheduler"
	"kurs_scheduler/pkg/utils"
	"strconv"
)

type QuantumCommand struct {
	Scheduler    *scheduler.Scheduler
	ErrorMessage string
}

func NewQuantumCommand(scheduler *scheduler.Scheduler) *QuantumCommand {
	return &QuantumCommand{
		Scheduler:    scheduler,
		ErrorMessage: "Использование: quantum <value>",
	}
}

func (cmd *QuantumCommand) Execute(args []string) {
	if len(args) != 1 {
		utils.Error(cmd.ErrorMessage)
		return
	}
	var quantum int
	var err error
	quantum, err = strconv.Atoi(args[0])
	if err != nil || quantum < 1 || quantum > 250 {
		utils.Error(cmd.ErrorMessage)
		return
	}

	cmd.Scheduler.SetQuantum(quantum)
	fmt.Println("Значение кванта времени успешно установлено")
}
