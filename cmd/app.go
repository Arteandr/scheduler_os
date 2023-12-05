package main

import (
	"fmt"
	"kurs_scheduler/internal/commands"
	"kurs_scheduler/internal/terminal"
)

func main() {
	// Канал для чтения cmd с терминала
	cmd := make(chan *commands.CommandArgs)
	term := terminal.NewTerminal(cmd)
	go term.Run()

	for {
		select {
		case args := <-cmd:
			fmt.Println(args)
			break
		}
	}
}
