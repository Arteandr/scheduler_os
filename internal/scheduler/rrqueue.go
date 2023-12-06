package scheduler

import (
	"fmt"
	"kurs_scheduler/internal/process"
	"kurs_scheduler/pkg/fifo"
	"math"
)

type RRQueue struct {
	queue        *fifo.Queue[process.Process]
	queueQuantum int
	quantum      *int
}

func NewRRQueue(quantum *int, queueQuantum int) *RRQueue {
	return &RRQueue{
		queue:        fifo.NewFIFOQueue[process.Process](),
		queueQuantum: queueQuantum,
		quantum:      quantum,
	}
}

func (q *RRQueue) Run(onTick OnTickCallback, enqueueToNextCallback EnqueueToNextCallback) {
	currentQueueQuantum := q.queueQuantum
	for currentQueueQuantum > 0 && q.queue.Size() > 0 {
		proc, err := q.queue.Dequeue()
		if err != nil {
			fmt.Println(err)
			continue
		}

		proc.Status = process.Running
		for i := int(math.Min(float64(*q.quantum), float64(proc.RemainingTime))); i != 0 && currentQueueQuantum > 0; i-- {
			currentQueueQuantum -= 1
			proc.RemainingTime -= 1
			onTick()
		}
		if int(proc.RemainingTime) > 0 {
			//proc.Priority += 1
			proc.Status = process.Waiting
			enqueueToNextCallback(proc)
		} else {
			proc.Status = process.Completed
		}
	}
}

func (q *RRQueue) Enqueue(item *process.Process) {
	q.queue.Enqueue(item)
}

func (q *RRQueue) Size() int {
	return q.queue.Size()
}
