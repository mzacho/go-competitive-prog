package queue

import "fmt"

type Queue[T interface{}] struct {
	q    []*T
	size int
	tail int
	head int
	full bool
}

func NewQueue[T interface{}](size int) *Queue[T] {
	return &Queue[T]{
		q:    make([]*T, size),
		size: size,
	}
}

func (q *Queue[T]) Enqueue(x *T) error {
	if q.full && q.head == q.tail {
		return fmt.Errorf("queue overflow")
	}
	q.q[q.tail] = x
	if q.tail == q.size-1 {
		q.tail = 0
	} else {
		q.tail += 1
	}
	q.full = q.head == q.tail
	return nil
}

func (q *Queue[T]) Dequeue() *T {
	if !q.full && q.head == q.tail {
		return nil
	}
	x := q.q[q.head]
	if q.head == q.size-1 {
		q.head = 0
	} else {
		q.head += 1
	}
	return x
}
