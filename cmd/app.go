package main

import (
	"fmt"
	"kurs_scheduler/internal/command"
	"kurs_scheduler/internal/scheduler"
	"kurs_scheduler/internal/terminal"
	"kurs_scheduler/pkg/utils"
)

func main() {
	utils.ClearScreen()
	// Канал для чтения cmd с терминала
	cmd := make(chan *command.Args)
	// Канал для ожидания выполнения CMD
	wait := make(chan struct{})
	term := terminal.NewTerminal(wait, cmd)
	go term.Run()

	processScheduler := scheduler.NewScheduler()
	cmdExecutor := command.NewCommandExecutor(processScheduler)

	for {
		select {
		case cmd := <-cmd:
			status := cmdExecutor.ExecuteCMD(cmd)
			if status == command.ErrorCmd {
				fmt.Printf("Команда \"%s\" не найдена\n", cmd.Alias)
			}
			wait <- struct{}{}
			break
		}
	}
}
