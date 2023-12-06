package command

import (
	"fmt"
	"kurs_scheduler/internal/scheduler"
	"strconv"
)

type BurstCommand struct {
	Scheduler    *scheduler.Scheduler
	ErrorMessage string
}

func NewBurstCommand(scheduler *scheduler.Scheduler) *BurstCommand {
	return &BurstCommand{
		Scheduler:    scheduler,
		ErrorMessage: "Использование: burst <value>",
	}
}

func (cmd *BurstCommand) Execute(args []string) {
	if len(args) != 1 {
		fmt.Println(cmd.ErrorMessage)
		return
	}
	var burst int
	var err error
	burst, err = strconv.Atoi(args[0])
	if err != nil || burst < 1 || burst > 250 {
		fmt.Println(cmd.ErrorMessage)
		return
	}

	cmd.Scheduler.SetMaxBurst(burst)
	fmt.Println("Значение CPU burst успешно установлено")
}
