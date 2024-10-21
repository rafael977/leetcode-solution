package designcirculardeque

/*
 * @lc app=leetcode id=641 lang=golang
 *
 * [641] Design Circular Deque
 */

// @lc code=start
type MyCircularDeque struct {
	cap  int
	list []int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		cap:  k,
		list: make([]int, 0, k),
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.list = append(this.list, 0)
	copy(this.list[1:], this.list)
	this.list[0] = value
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.list = append(this.list, value)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.list = this.list[1:]
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.list = this.list[:len(this.list)-1]
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.list[0]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.list[len(this.list)-1]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return len(this.list) == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return len(this.list) == this.cap
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
// @lc code=end
