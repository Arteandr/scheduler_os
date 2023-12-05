package scheduler

import "kurs_scheduler/internal/process"

type Scheduler struct {
	Quantum   uint8
	MaxBurst  uint8
	Processes []*process.Process
}

func NewScheduler() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) SetQuantum(quantum uint8) {
	s.Quantum = quantum
}

func (s *Scheduler) GenerateProcesses(count int) {
	if count < 2 {
		return
	}

	s.Processes = make([]*process.Process, count)
	for i := 0; i < count; i++ {
		s.Processes = append(s.Processes, process.GenerateProcess(i, int(s.MaxBurst)))
	}
}

func (s *Scheduler) GetAllProcesses() *[]*process.Process {
	return &s.Processes
}
