package leetcodesolution

/*
 * @lc app=leetcode id=786 lang=golang
 *
 * [786] K-th Smallest Prime Fraction
 */

// @lc code=start

import "github.com/emirpasic/gods/v2/queues/priorityqueue"

func kthSmallestPrimeFraction(arr []int, k int) []int {
	type frac struct {
		i, j int
	}

	compareFrac := func(f1, f2 *frac) int {
		return arr[f1.i]*arr[f2.j] - arr[f2.i]*arr[f1.j]
	}

	pq := priorityqueue.NewWith(compareFrac)
	for j := 1; j < len(arr); j++ {
		pq.Enqueue(&frac{i: 0, j: j})
	}

	for pq.Size() > 0 {
		x, _ := pq.Dequeue()
		k--
		if k == 0 {
			return []int{arr[x.i], arr[x.j]}
		}
		if x.i+1 < x.j {
			pq.Enqueue(&frac{i: x.i + 1, j: x.j})
		}
	}

	return []int{}
}

// @lc code=end
