package scheduler

import "kurs_scheduler/internal/process"

type OnTickCallback = func()
type EnqueueToNextCallback = func(*process.Process)

type ISchedulerQueue interface {
	Run(onTickCallback OnTickCallback,
		enqueueToNextCallback EnqueueToNextCallback)
	Enqueue(item *process.Process)
	Size() int
}
