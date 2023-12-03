package seatreservationmanager

import (
	"container/heap"
	"sort"
)

/*
 * @lc app=leetcode id=1845 lang=golang
 *
 * [1845] Seat Reservation Manager
 */

// @lc code=start
type IntHeap struct {
	sort.IntSlice
}

func (h *IntHeap) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *IntHeap) Pop() any {
	old := h.IntSlice
	n := len(old)
	x := old[n-1]
	h.IntSlice = old[0 : n-1]
	return x
}

type SeatManager struct {
	aval *IntHeap
}

func Constructor(n int) SeatManager {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = i + 1
	}
	h := &IntHeap{
		sort.IntSlice(arr),
	}
	heap.Init(h)
	return SeatManager{aval: h}
}

func (this *SeatManager) Reserve() int {
	return heap.Pop(this.aval).(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(this.aval, seatNumber)
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
// @lc code=end
