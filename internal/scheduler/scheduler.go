package scheduler

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"kurs_scheduler/internal/process"
	"kurs_scheduler/pkg/utils"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
)

type Scheduler struct {
	Quantum         int
	MaxBurst        int
	Processes       []*process.Process
	ProcessHistory  [][]process.Process
	CurrentTick     int
	MultilevelQueue *MultilevelQueue
}

func NewScheduler() *Scheduler {
	sched := &Scheduler{
		CurrentTick: 0,
	}
	sched.MultilevelQueue = NewMultilevelQueue(sched)

	return sched
}

func (s *Scheduler) Run() {
	s.SetMaxBurst(10)
	s.GenerateProcesses(10)
	if len(s.Processes) < 1 {
		return
	}

	for s.MultilevelQueue.TotalSize() > 0 {
		for i, queue := range s.MultilevelQueue.queues {
			if queue.Size() < 1 {
				continue
			}

			enqueueToNext := func(s *Scheduler, currentQueue int) func(process *process.Process) {
				return func(process *process.Process) {
					process.Priority += 1
					s.MultilevelQueue.queues[process.Priority].Enqueue(process)
				}
			}(s, i)

			onTickCallback := func(s *Scheduler) func() {
				return func() {
					s.SaveHistory()
				}
			}(s)

			queue.Run(onTickCallback, enqueueToNext)
		}
	}
	s.SaveHistory()
}

func (s *Scheduler) SaveHistory() {
	s.ProcessHistory = append(s.ProcessHistory, utils.PtrToSlice(&s.Processes))
}

func (s *Scheduler) RunDraw() {
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
		case char == 'n':
			if s.CurrentTick != len(s.ProcessHistory)-1 {
				s.CurrentTick += 1
			}
			break
		case char == 'p':
			if s.CurrentTick != 0 {
				s.CurrentTick -= 1
			}
			break
		default:
			continue
		}
	}
}

func (s *Scheduler) draw() {
	switch runtime.GOOS {
	case "windows":
		clear := exec.Command("cmd", "/c", "cls")
		clear.Stdout = os.Stdout
		clear.Run()
		break
	case "linux":
		clear := exec.Command("clear")
		clear.Stdout = os.Stdout
		clear.Run()
		break
	}

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
	t.Style().Options.SeparateRows = true
	t.AppendHeader(s.generateHeader())
	t.AppendRows(s.generateRows())
	t.AppendSeparator()
	t.Render()
}

func (s *Scheduler) generateHeader() table.Row {
	green := color.New(color.FgGreen).SprintfFunc()
	headerRow := table.Row{green(strconv.Itoa(s.CurrentTick))}
	for i := 1; i <= s.getTotalBurst()+1; i++ {
		headerRow = append(headerRow, green(strconv.Itoa(i)))
	}

	return headerRow
}

func (s *Scheduler) generateRows() []table.Row {
	green := color.New(color.FgGreen, color.Underline).SprintfFunc()
	t := make([]table.Row, len(s.Processes))
	for i, proc := range s.ProcessHistory[s.CurrentTick] {
		info := fmt.Sprintf("%s\nUID: %d\nCPU Burst: %d Ост.время: %d\n"+
			"Приоритет: %d Статус: %d",
			green(fmt.Sprintf("%s: %s", "PID", color.New(color.Bold).SprintFunc()(proc.ID))),
			proc.UID,
			proc.Burst,
			proc.RemainingTime,
			proc.Priority,
			proc.Status,
		)
		t[i] = table.Row{info}
		for j := 0; j <= s.CurrentTick; j++ {
			t[i] = append(t[i], s.ProcessHistory[j][proc.ID].StringStatus())
		}
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
		s.MultilevelQueue.queues[0].Enqueue(proc)
	}
}

func (s *Scheduler) GetAllProcesses() []*process.Process {
	return s.Processes
}
