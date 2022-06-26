package coursescheduleiii

import (
	"container/heap"
	"sort"
)

/*
 * @lc app=leetcode id=630 lang=golang
 *
 * [630] Course Schedule III
 */

// @lc code=start
func scheduleCourse(courses [][]int) (ans int) {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	day := 0
	cnt := 0
	ih := &IntHeap{}
	heap.Init(ih)
	for i := range courses {
		day += courses[i][0]
		cnt++
		heap.Push(ih, courses[i][0])
		if day <= courses[i][1] {
			ans = max(ans, cnt)
		} else {
			for day > courses[i][1] {
				day -= heap.Pop(ih).(int)
				cnt--
			}
		}
	}
	return
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
