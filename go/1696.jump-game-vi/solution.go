package jumpgamevi

import "container/heap"

/*
 * @lc app=leetcode id=1696 lang=golang
 *
 * [1696] Jump Game VI
 */

// @lc code=start
type nv struct {
	idx, v int
}

type nvHeap []nv

func maxResult(nums []int, k int) (ans int) {
	n := len(nums)
	ans = nums[0]

	nvh := nvHeap{nv{idx: 0, v: ans}}
	heap.Init(&nvh)

	for i := 1; i < n; i++ {
		for i-nvh[0].idx > k {
			heap.Pop(&nvh)
		}

		ans = nvh[0].v + nums[i]
		heap.Push(&nvh, nv{idx: i, v: ans})
	}
	return
}

func (h nvHeap) Len() int           { return len(h) }
func (h nvHeap) Less(i, j int) bool { return h[i].v > h[j].v }
func (h nvHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *nvHeap) Push(x interface{}) {
	*h = append(*h, x.(nv))
}

func (h *nvHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// @lc code=end
