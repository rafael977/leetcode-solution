package main

/*
 * @lc app=leetcode id=895 lang=golang
 *
 * [895] Maximum Frequency Stack
 */

// @lc code=start
type FreqStack struct {
	freq  map[int]int
	group map[int][]int
	max   int
}

func Constructor() FreqStack {
	return FreqStack{
		freq:  make(map[int]int),
		group: make(map[int][]int),
		max:   0,
	}
}

func (this *FreqStack) Push(val int) {
	this.freq[val]++
	if this.freq[val] > this.max {
		this.max = this.freq[val]
	}
	if this.group[this.freq[val]] == nil {
		this.group[this.freq[val]] = make([]int, 0)
	}
	this.group[this.freq[val]] = append(this.group[this.freq[val]], val)
}

func (this *FreqStack) Pop() int {
	mg := this.group[this.max]
	val := mg[len(mg)-1]
	this.freq[val]--
	this.group[this.max] = mg[:len(mg)-1]
	if len(this.group[this.max]) == 0 {
		this.max--
	}
	return val
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
// @lc code=end
