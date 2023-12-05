package terminal

import (
	"bufio"
	"fmt"
	"kurs_scheduler/internal/commands"
	"os"
	"strings"
)

type AppTerminal struct {
	OutputChannel chan<- *commands.CommandArgs
	InputChannel  <-chan struct{}
}

func NewTerminal(
	inputChannel <-chan struct{},
	outputChannel chan<- *commands.CommandArgs) *AppTerminal {
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
		t.OutputChannel <- commands.NewCommandArgs(line[0], line[1:])
		<-t.InputChannel
	}
}

func (t *AppTerminal) GetLine() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("scheduler> ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода")
		return ""
	}

	return strings.TrimSpace(text)
}
