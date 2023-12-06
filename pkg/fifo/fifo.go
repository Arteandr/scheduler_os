package fifo

import (
	"errors"
)

type Queue struct {
	items []interface{}
}

func NewFIFOQueue() *Queue {
	return &Queue{items: []interface{}{}}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (interface{}, error) {
	if len(q.items) == 0 {
		return nil, errors.New("очередь пустая")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func (q *Queue) Front() (interface{}, error) {
	if len(q.items) == 0 {
		return nil, errors.New("очередь пустая")
	}
	return q.items[0], nil
}

func (q *Queue) Size() int {
	return len(q.items)
}
