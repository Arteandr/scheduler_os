package process

import (
	"math/rand"
)

type Status = uint8

const (
	Running   Status = iota // Выполнение
	Readiness               // Готовность
	Waiting                 // Ожидание
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
	return &Process{
		ID:     uint32(id),
		UID:    uint32(rand.Intn(10)),
		Burst:  uint32(rand.Intn(maxBurst)),
		Status: Readiness,
	}
}
