package command

import (
	"fmt"
	"kurs_scheduler/internal/scheduler"
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
		fmt.Println(cmd.ErrorMessage)
		return
	}
	var quantum uint64
	var err error
	quantum, err = strconv.ParseUint(args[0], 10, 8)
	if err != nil || quantum < 1 {
		fmt.Println(cmd.ErrorMessage)
		return
	}

	cmd.Scheduler.SetQuantum(uint8(quantum))
	fmt.Println("Значение кванта времени успешно установлено")
}
