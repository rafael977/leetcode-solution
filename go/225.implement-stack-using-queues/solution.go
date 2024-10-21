package implementstackusingqueues

/*
 * @lc app=leetcode id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 */

// @lc code=start
type Queue struct {
	arr []int
}

func (q Queue) Size() int {
	return len(q.arr)
}

func (q Queue) IsEmpty() bool {
	return len(q.arr) == 0
}

func (q *Queue) Enqueue(v int) {
	q.arr = append(q.arr, v)
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		return 0
	}
	v := q.arr[0]
	q.arr = q.arr[1:]
	return v
}

func (q Queue) Peek() int {
	return q.arr[0]
}

type MyStack struct {
	q *Queue
}

func Constructor() MyStack {
	return MyStack{q: &Queue{}}
}

func (this *MyStack) Push(x int) {
	n := this.q.Size()
	this.q.Enqueue(x)
	for i := 0; i < n; i++ {
		this.q.Enqueue(this.q.Dequeue())
	}
}

func (this *MyStack) Pop() int {
	return this.q.Dequeue()
}

func (this *MyStack) Top() int {
	return this.q.Peek()
}

func (this *MyStack) Empty() bool {
	return this.q.IsEmpty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end
