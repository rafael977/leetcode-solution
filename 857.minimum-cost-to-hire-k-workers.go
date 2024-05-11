package leetcodesolution

/*
 * @lc app=leetcode id=857 lang=golang
 *
 * [857] Minimum Cost to Hire K Workers
 */

// @lc code=start
import (
	"sort"

	"github.com/emirpasic/gods/v2/queues/priorityqueue"
)

func mincostToHireWorkers(quality []int, wage []int, k int) (ans float64) {
	type tuple struct {
		q, w int
	}
	workers := make([]*tuple, 0, len(quality))
	for i, q := range quality {
		workers = append(workers, &tuple{q: q, w: wage[i]})
	}

	sort.Slice(workers, func(i, j int) bool {
		return workers[i].w*workers[j].q < workers[j].w*workers[i].q
	})

	pq := priorityqueue.NewWith(func(t1, t2 *tuple) int { return t2.q - t1.q })
	sumq := 0
	for _, v := range workers[:k] {
		sumq += v.q
		pq.Enqueue(v)
	}
	ans = float64(sumq*workers[k-1].w) / float64(workers[k-1].q)

	for _, v := range workers[k:] {
		top, _ := pq.Peek()
		if v.q < top.q {
			pq.Dequeue()
			pq.Enqueue(v)
			sumq += -top.q + v.q
			ans = min(ans, float64(sumq*v.w)/float64(v.q))
		}
	}

	return
}

// @lc code=end
