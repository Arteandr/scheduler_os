package terminal

import (
	"bufio"
	"fmt"
	"kurs_scheduler/internal/command"
	"kurs_scheduler/pkg/utils"
	"os"
	"strings"
)

type AppTerminal struct {
	OutputChannel chan<- *command.Args
	InputChannel  <-chan struct{}
}

func NewTerminal(
	inputChannel <-chan struct{},
	outputChannel chan<- *command.Args) *AppTerminal {
	return &AppTerminal{
		OutputChannel: outputChannel,
		InputChannel:  inputChannel,
	}
}

func (t *AppTerminal) Run() {
	for {
		line := strings.Split(t.GetLine(), " ")
		if len(line) < 0 {
			continue
		}
		t.OutputChannel <- command.NewCommandArgs(line[0], line[1:])
		<-t.InputChannel
	}
}

func (t *AppTerminal) GetLine() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("scheduler> ")
	text, err := reader.ReadString('\n')
	if err != nil {
		utils.Error("Ошибка ввода")
		return ""
	}

	return strings.TrimSpace(text)
}
