package common

import "errors"

type Queue[T any] struct {
	items []T
}

func (queue *Queue[T]) Enqueue(value T) {
	queue.items = append(queue.items, value)
}

func (queue *Queue[T]) Dequeue() (any, error) {
	if len(queue.items) > 0 {
		res := queue.items[0]
		if len(queue.items) == 1 {
			queue.items = make([]T, 0)
		} else {
			queue.items = queue.items[1:]
		}
		return res, nil
	}
	return nil, errors.New("queue is empty")
}

func (queue *Queue[T]) Size() int {
	return len(queue.items)
}

func (queue *Queue[T]) Empty() bool {
	return queue.Size() == 0
}
