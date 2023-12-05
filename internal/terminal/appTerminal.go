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
}

func NewTerminal(outputChannel chan<- *commands.CommandArgs) *AppTerminal {
	return &AppTerminal{
		OutputChannel: outputChannel,
	}
}

func (t *AppTerminal) Run() {
	for {
		line := strings.Split(t.GetLine(), " ")
		if len(line) > 0 {
			t.OutputChannel <- commands.NewCommandArgs(line[0], line[1:])
		}
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

	return text
}
