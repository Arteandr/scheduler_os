package fifo

import (
	"errors"
)

type Queue[T any] struct {
	items []*T
}

func NewFIFOQueue[T comparable]() *Queue[T] {
	return &Queue[T]{items: []*T{}}
}

func (q *Queue[T]) Enqueue(item *T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (*T, error) {
	if len(q.items) == 0 {
		return nil, errors.New("очередь пустая")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func (q *Queue[T]) Front() (*T, error) {
	if len(q.items) == 0 {
		return nil, errors.New("очередь пустая")
	}
	return q.items[0], nil
}

func (q *Queue[T]) Size() int {
	return len(q.items)
}
