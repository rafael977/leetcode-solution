package furthestbuildingyoucanreach

import "container/heap"

/*
 * @lc app=leetcode id=1642 lang=golang
 *
 * [1642] Furthest Building You Can Reach
 */

// @lc code=start
func furthestBuilding(heights []int, bricks int, ladders int) int {
	hdiff := make([]int, len(heights))
	for i := range heights {
		if i == 0 {
			continue
		}
		if heights[i-1] >= heights[i] {
			hdiff[i] = 0
		} else {
			hdiff[i] = heights[i] - heights[i-1]
		}
	}

	i := 0
	ih := &IntHeap{}
	heap.Init(ih)
	sum := 0
	for i < len(hdiff) {
		sum += hdiff[i]
		heap.Push(ih, hdiff[i])
		for sum > bricks && ladders > 0 {
			sum -= heap.Pop(ih).(int)
			ladders--
		}
		if sum > bricks {
			break
		}
		i++
	}
	return i - 1
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// @lc code=end
