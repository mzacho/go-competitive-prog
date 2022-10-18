package queue_test

import (
	"competitive_programming/util/queue"
	"testing"
)

func TestQueueUnderflow(t *testing.T) {
	q := queue.NewQueue[string](4)
	x := q.Dequeue()
	if x != nil {
		t.Errorf("q.Dequeue() = %v; want nil", x)
	}
}

func TestQueueOverflow(t *testing.T) {
	q := queue.NewQueue[string](2)
	a := "a"
	b := "b"
	c := "c"
	err := q.Enqueue(&a)
	if err != nil {
		t.Errorf("q.Enqueue() = %v; expected nil err", err)
	}
	err = q.Enqueue(&b)
	if err != nil {
		t.Errorf("q.Enqueue() = %v; expected nil err", err)
	}
	err = q.Enqueue(&c)
	if err == nil {
		t.Errorf("q.Enqueue() = %v; expected overflow err", err)
	}
}
