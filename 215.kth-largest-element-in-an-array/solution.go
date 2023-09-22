package kthlargestelementinanarray

import (
	"container/heap"
	"sort"
)

/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start

// func findKthLargest(nums []int, k int) int {
// 	sort.Ints(nums)
// 	return nums[len(nums)-k]
// }

type IntHeap struct {
	sort.IntSlice
}

func (h *IntHeap) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *IntHeap) Pop() any {
	old := h.IntSlice
	n := h.Len()
	x := old[n-1]
	h.IntSlice = old[:n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)

	for _, v := range nums {
		heap.Push(h, v)
		for h.Len() > k {
			heap.Pop(h)
		}
	}

	return heap.Pop(h).(int)
}

// @lc code=end
