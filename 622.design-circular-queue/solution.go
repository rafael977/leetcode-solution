package designcircularqueue

/*
 * @lc app=leetcode id=622 lang=golang
 *
 * [622] Design Circular Queue
 */

// @lc code=start
type MyCircularQueue struct {
	q          []int
	cap        int
	head, tail int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		q:    make([]int, k),
		cap:  k,
		head: 0,
		tail: 0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	size := this.tail - this.head
	if size == this.cap {
		return false
	}
	this.q[this.tail%this.cap] = value
	this.tail++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.head < this.tail {
		this.head++
		return true
	}
	return false
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	idx := this.head % this.cap
	return this.q[idx]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	idx := (this.tail - 1) % this.cap
	return this.q[idx]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == this.tail
}

func (this *MyCircularQueue) IsFull() bool {
	return this.tail-this.head == this.cap
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
// @lc code=end
