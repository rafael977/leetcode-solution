package smallestrangecoveringelementsfromklists

/*
 * @lc app=leetcode id=632 lang=golang
 *
 * [632] Smallest Range Covering Elements from K Lists
 */

// @lc code=start

import pq "github.com/emirpasic/gods/v2/queues/priorityqueue"

func smallestRange(nums [][]int) []int {
	inf := int(1e5 + 1)
	type El struct {
		i, j, val int
	}
	l, r, maxV := -inf, inf, -inf

	q := pq.NewWith(func(x, y El) int {
		return x.val - y.val
	})
	for i, n := range nums {
		v := n[0]
		q.Enqueue(El{i: i, j: 0, val: v})
		maxV = max(maxV, v)
	}
	for q.Size() == len(nums) {
		e, _ := q.Dequeue()
		if maxV-e.val < r-l {
			l = e.val
			r = maxV
		}

		if e.j+1 < len(nums[e.i]) {
			q.Enqueue(El{i: e.i, j: e.j + 1, val: nums[e.i][e.j+1]})
			maxV = max(maxV, nums[e.i][e.j+1])
		}
	}
	return []int{l, r}
}

// @lc code=end
