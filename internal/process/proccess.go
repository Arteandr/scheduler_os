package process

import (
	"github.com/fatih/color"
	"math/rand"
)

const (
	MinPriority = 0
	MaxPriority = 2
)

type Status = uint8

const (
	Waiting   Status = iota // Ожидание
	Running                 // Выполнение
	Completed               // Выполнен
)

type Process struct {
	ID            uint32
	UID           uint32
	Burst         uint32
	RemainingTime uint32
	ArrivalTime   uint32
	Status        Status
	Priority      uint32
}

func GenerateProcess(id, maxBurst int) *Process {
	p := &Process{
		ID:       uint32(id),
		UID:      uint32(rand.Intn(10)),
		Burst:    uint32(rand.Intn(maxBurst)) + 1,
		Priority: uint32(MinPriority),
		Status:   Waiting,
	}
	p.RemainingTime = p.Burst

	return p
}

func (proc *Process) StringStatus() string {
	cyan := color.New(color.FgHiCyan).SprintfFunc()
	red := color.New(color.FgHiRed).SprintfFunc()
	green := color.New(color.FgHiGreen).SprintfFunc()
	if proc.Status == Waiting {
		return cyan("W")
	} else if proc.Status == Running {
		return red("R")
	} else {
		return green("F")
	}
}

func (proc *Process) LongStringStatus() string {
	cyan := color.New(color.FgHiCyan).SprintfFunc()
	red := color.New(color.FgHiRed).SprintfFunc()
	green := color.New(color.FgHiGreen).SprintfFunc()
	if proc.Status == Waiting {
		return cyan("готовность")
	} else if proc.Status == Running {
		return red("выполнение")
	} else {
		return green("выполнен")
	}
}
