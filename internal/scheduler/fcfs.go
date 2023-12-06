package scheduler

import (
	"fmt"
	"kurs_scheduler/internal/process"
	"kurs_scheduler/pkg/fifo"
)

type FCFSQueue struct {
	queue *fifo.Queue[process.Process]
}

func NewFCRSQueue() *FCFSQueue {
	return &FCFSQueue{
		queue: fifo.NewFIFOQueue[process.Process](),
	}
}

func (q *FCFSQueue) Run(onTickCallback OnTickCallback) {
	for q.Size() > 0 {
		proc, err := q.queue.Dequeue()
		if err != nil {
			fmt.Println(err)
			continue
		}

		for int(proc.RemainingTime) > 0 {
			proc.RemainingTime -= 1
			onTickCallback()
		}

		proc.Status = process.Completed
	}
}

func (q *FCFSQueue) Enqueue(item *process.Process) {
	q.queue.Enqueue(item)
}

func (q *FCFSQueue) Size() int {
	return q.queue.Size()
}
