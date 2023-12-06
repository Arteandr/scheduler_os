package scheduler

import "kurs_scheduler/internal/process"

type OnTickCallback = func()

type ISchedulerQueue interface {
	//Run(onTick chan<- struct{})
	//Run(onTick chan<- int)
	Run(onTickCallback OnTickCallback)
	Enqueue(item *process.Process)
	Size() int
}
