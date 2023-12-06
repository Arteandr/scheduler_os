package scheduler

type MultilevelQueue struct {
	queues []ISchedulerQueue
}

func NewMultilevelQueue() *MultilevelQueue {
	mq := &MultilevelQueue{}
	mq.queues = append(mq.queues, NewRRQueue())
	mq.queues = append(mq.queues, NewRRQueue())
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
