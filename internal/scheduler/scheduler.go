package scheduler

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"kurs_scheduler/internal/process"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
)

type Scheduler struct {
	Quantum     int
	MaxBurst    int
	Processes   []*process.Process
	CurrentTick int
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		CurrentTick: 0,
	}
}

func (s *Scheduler) Run() {
	if err := keyboard.Open(); err != nil {
		fmt.Println("Невозможно запустить визуализацию")
		return
	}
	defer keyboard.Close()

	quit := make(chan struct{})
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig
		close(quit)
	}()

main:
	for {
		s.draw()

		char, key, err := keyboard.GetKey()
		if err != nil {
			return
		}
		switch {
		case char == 'q' || key == keyboard.KeyCtrlC:
			break main
		default:
			continue
		}
	}
}

func (s *Scheduler) draw() {
	clear := exec.Command("cmd", "/c", "cls")
	clear.Stdout = os.Stdout
	clear.Run()

	s.drawTable()
	s.drawHotkeys()
}

func (s *Scheduler) getTotalBurst() int {
	sum := 0
	for _, proc := range s.Processes {
		sum += int(proc.Burst)
	}

	return sum
}

func (s *Scheduler) drawTable() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.AppendHeader(s.generateHeader())

	t.Style().Options.SeparateRows = true

	t.AppendRows(s.generateRows())
	t.AppendSeparator()
	t.Render()
}

func (s *Scheduler) generateHeader() table.Row {
	green := color.New(color.FgGreen).SprintfFunc()
	headerRow := table.Row{green("PID")}
	for i := 1; i <= s.getTotalBurst(); i++ {
		headerRow = append(headerRow, strconv.Itoa(i))
	}

	return headerRow
}

func (s *Scheduler) generateRows() []table.Row {
	t := make([]table.Row, len(s.Processes))
	for i, proc := range s.Processes {
		info := fmt.Sprintf("PID: %d UID: %d Burst: %d\n"+
			"Приоритет: %d Статус: %d", proc.ID, proc.UID, proc.Burst, proc.Status, proc.Priority)
		t[i] = table.Row{info}
	}

	return t
}

func (s *Scheduler) drawHotkeys() {
	red := color.New(color.FgRed).SprintfFunc()
	green := color.New(color.FgGreen).SprintfFunc()
	fmt.Printf("\n%s Выход\t%s Предыдущий тик\t%s Следующий тик\n",
		red("(q)"),
		green("(p)"),
		green("(n)"))

}

func (s *Scheduler) NextTick() {
	if len(s.Processes) < 1 {
		return
	}

	s.CurrentTick += 1
}

func (s *Scheduler) PrevTick() {
	if s.CurrentTick == 1 || len(s.Processes) < 1 {
		return
	}

	s.CurrentTick -= 1
}

func (s *Scheduler) SetQuantum(quantum int) {
	if quantum < 1 {
		return
	}

	s.Quantum = quantum
}

func (s *Scheduler) SetMaxBurst(burst int) {
	if burst < 1 {
		return
	}

	s.MaxBurst = burst
}

func (s *Scheduler) GenerateProcesses(count int) {
	if count < 2 {
		return
	}

	s.Processes = make([]*process.Process, count)
	for i := 0; i < count; i++ {
		proc := process.GenerateProcess(i, s.MaxBurst)
		s.Processes[i] = proc
	}
}

func (s *Scheduler) GetAllProcesses() []*process.Process {
	return s.Processes
}
