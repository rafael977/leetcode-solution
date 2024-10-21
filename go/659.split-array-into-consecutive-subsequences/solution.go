package splitarrayintoconsecutivesubsequences

import (
	"container/heap"
	"sort"
)

/*
 * @lc app=leetcode id=659 lang=golang
 *
 * [659] Split Array into Consecutive Subsequences
 */

// @lc code=start
type minHeap struct {
	sort.IntSlice
}

func (h *minHeap) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *minHeap) Pop() interface{} {
	v := h.IntSlice[h.Len()-1]
	h.IntSlice = h.IntSlice[:h.Len()-1]
	return v
}

func (h *minHeap) push(v int) {
	heap.Push(h, v)
}

func (h *minHeap) pop() int {
	return heap.Pop(h).(int)
}

func isPossible(nums []int) bool {
	lm := make(map[int]*minHeap)
	for _, v := range nums {
		if lm[v] == nil {
			lm[v] = new(minHeap)
		}
		if lh := lm[v-1]; lh != nil {
			lm[v].push(lh.pop() + 1)
			if lh.Len() == 0 {
				delete(lm, v-1)
			}
		} else {
			lm[v].push(1)
		}
	}

	for _, l := range lm {
		if l.IntSlice[0] < 3 {
			return false
		}
	}
	return true
}

// @lc code=end
