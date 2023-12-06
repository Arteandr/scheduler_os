package process

import (
	"math/rand"
)

const (
	MinPriority = 0
	MaxPriority = 2
)

type Status = uint8

const (
	Running   Status = iota // Выполнение
	Readiness               // Готовность
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
		Status:   Readiness,
	}
	p.RemainingTime = p.Burst

	return p
}

func (proc *Process) IncreasePriority() {
	proc.Priority += 1
	if proc.Priority > MaxPriority {
		proc.Priority = MinPriority
	}
}
