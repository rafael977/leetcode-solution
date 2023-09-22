package slidingwindowmaximum

import "container/heap"

/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */

// @lc code=start
type NumHeap struct {
	Nums []int
	Idxs []int
}

func (h NumHeap) Len() int {
	return len(h.Idxs)
}

func (h NumHeap) Less(i, j int) bool {
	return h.Nums[h.Idxs[i]] > h.Nums[h.Idxs[j]]
}

func (h NumHeap) Swap(i, j int) {
	h.Idxs[i], h.Idxs[j] = h.Idxs[j], h.Idxs[i]
}

func (h *NumHeap) Push(x any) {
	h.Idxs = append(h.Idxs, x.(int))
}

func (h *NumHeap) Pop() any {
	old := h.Idxs
	n := len(old)
	x := old[n-1]
	h.Idxs = old[:n-1]
	return x
}

func maxSlidingWindow(nums []int, k int) (ans []int) {
	h := NumHeap{
		Nums: nums,
		Idxs: make([]int, 0, k),
	}
	heap.Init(&h)

	for i := 0; i < k; i++ {
		heap.Push(&h, i)
	}
	ans = append(ans, nums[h.Idxs[0]])
	for i := k; i < len(nums); i++ {
		heap.Push(&h, i)
		for h.Idxs[0] <= i-k {
			heap.Pop(&h)
		}
		ans = append(ans, nums[h.Idxs[0]])
	}

	return
}

// @lc code=end
