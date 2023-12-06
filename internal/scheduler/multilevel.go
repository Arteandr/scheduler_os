package scheduler

const (
	FirstQueueQuantum  = 8
	SecondQueueQuantum = FirstQueueQuantum * 2
)

type MultilevelQueue struct {
	queues []ISchedulerQueue
}

func NewMultilevelQueue(s *Scheduler) *MultilevelQueue {
	mq := &MultilevelQueue{}
	mq.queues = append(mq.queues, NewRRQueue(&s.Quantum, FirstQueueQuantum))
	mq.queues = append(mq.queues, NewRRQueue(&s.Quantum, SecondQueueQuantum))
	mq.queues = append(mq.queues, NewFCRSQueue())

	return mq
}

func (mq *MultilevelQueue) Queue(index int) ISchedulerQueue {
	return mq.queues[index]
}

func (mq *MultilevelQueue) TotalSize() int {
	size := 0
	for _, queue := range mq.queues {
		size += queue.Size()
	}

	return size
}
