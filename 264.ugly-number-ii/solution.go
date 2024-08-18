package uglynumberii

/*
 * @lc app=leetcode id=264 lang=golang
 *
 * [264] Ugly Number II
 */

// @lc code=start
import pq "github.com/emirpasic/gods/v2/queues/priorityqueue"

func nthUglyNumber(n int) int {
	c := 0
	seen := make(map[int]bool)
	q := pq.New[int]()
	q.Enqueue(1)
	seen[1] = true
	for {
		v, _ := q.Dequeue()
		c++
		if c == n {
			return v
		}
		if !seen[v*2] {
			q.Enqueue(v * 2)
			seen[v*2] = true
		}
		if !seen[v*3] {
			q.Enqueue(v * 3)
			seen[v*3] = true
		}
		if !seen[v*5] {
			q.Enqueue(v * 5)
			seen[v*5] = true
		}
	}
}

// @lc code=end
