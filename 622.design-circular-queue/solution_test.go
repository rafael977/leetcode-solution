package designcircularqueue

import "testing"

func Test_MyCircularQueue(t *testing.T) {
	q := Constructor(3)
	assert(t, "enQueue", q.EnQueue(1), true)
	assert(t, "enQueue", q.EnQueue(2), true)
	assert(t, "enQueue", q.EnQueue(3), true)
	assert(t, "enQueue", q.EnQueue(4), false)
	assert(t, "rear", q.Rear(), 3)
	assert(t, "isFull", q.IsFull(), true)
	assert(t, "deQueue", q.DeQueue(), true)
	assert(t, "enQueue", q.EnQueue(4), true)
	assert(t, "rear", q.Rear(), 4)
}

func assert[T int | bool](t *testing.T, f string, actual, expect T) {
	if actual != expect {
		t.Errorf("%s: expect %v, got %v", f, expect, actual)
	}
}
