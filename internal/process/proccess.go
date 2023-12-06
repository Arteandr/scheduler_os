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
	Waiting   Status = iota // Готовность
	Running                 // Выполнение
	Completed               // Выполнена
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
	yellow := color.New(color.FgHiYellow).SprintlnFunc()
	red := color.New(color.FgHiRed).SprintlnFunc()
	if proc.Status == Waiting {
		return yellow("W")
	} else if proc.Status == Running {
		return red("R")
	} else {
		return cyan("F")
	}
}
