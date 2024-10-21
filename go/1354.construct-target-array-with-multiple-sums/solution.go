package constructtargetarraywithmultiplesums

import "container/heap"

/*
 * @lc app=leetcode id=1354 lang=golang
 *
 * [1354] Construct Target Array With Multiple Sums
 */

// @lc code=start
func isPossible(target []int) bool {
	if len(target) == 1 {
		return target[0] == 1
	}

	sum := 0
	for _, v := range target {
		sum += v
	}

	ih := IntHeap(target)
	heap.Init(&ih)
	for ih.Len() > 0 {
		s := heap.Pop(&ih).(int)
		if s == 1 {
			// max is 1
			break
		}

		nxtS := ih[0]
		sumRest := sum - s
		k := (s-nxtS)/sumRest + 1
		if nxtS == 1 {
			k = (s - nxtS + sumRest - 1) / sumRest
		}
		r := s - k*sumRest
		if r < 1 {
			return false
		}
		sum = sum - s + r
		heap.Push(&ih, r)
	}
	return true
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
