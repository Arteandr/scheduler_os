package main

import (
	"fmt"
	"kurs_scheduler/internal/commands"
	"kurs_scheduler/internal/terminal"
)

func main() {
	// Канал для чтения cmd с терминала
	cmd := make(chan *commands.CommandArgs)
	// Канал для ожидания выполнения CMD
	wait := make(chan struct{})
	term := terminal.NewTerminal(wait, cmd)
	go term.Run()

	cmdExecutor := commands.NewCommandExecutor()

	for {
		select {
		case cmd := <-cmd:
			status := cmdExecutor.ExecuteCMD(cmd)
			fmt.Println("Статус выполнения CMD: ", status)
			wait <- struct{}{}
			break
		}
	}
}
