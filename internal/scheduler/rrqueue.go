package scheduler

import (
	"kurs_scheduler/internal/process"
	"kurs_scheduler/pkg/fifo"
)

type RRQueue struct {
	queue *fifo.Queue[process.Process]
}

func NewRRQueue() *RRQueue {
	return &RRQueue{
		queue: fifo.NewFIFOQueue[process.Process](),
	}
}

func (q *RRQueue) Run(onTick OnTickCallback) {

}

func (q *RRQueue) Enqueue(item *process.Process) {
	q.queue.Enqueue(item)
}

func (q *RRQueue) Size() int {
	return q.queue.Size()
}
