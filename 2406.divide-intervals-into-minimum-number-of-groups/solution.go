package divideintervalsintominimumnumberofgroups

/*
 * @lc app=leetcode id=2406 lang=golang
 *
 * [2406] Divide Intervals Into Minimum Number of Groups
 */

// @lc code=start
import (
	"sort"

	pq "github.com/emirpasic/gods/v2/queues/priorityqueue"
)

func minGroups(intervals [][]int) (ans int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	q := pq.NewWith(func(x, y [2]int) int {
		return x[1] - y[1]
	})

	for _, v := range intervals {
		for !q.Empty() {
			if h, _ := q.Peek(); h[1] < v[0] {
				q.Dequeue()
			} else {
				break
			}
		}
		q.Enqueue([2]int{v[0], v[1]})
		ans = max(ans, q.Size())
	}

	return
}

// @lc code=end
