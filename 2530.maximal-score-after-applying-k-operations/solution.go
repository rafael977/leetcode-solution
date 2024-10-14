package maximalscoreafterapplyingkoperations

/*
 * @lc app=leetcode id=2530 lang=golang
 *
 * [2530] Maximal Score After Applying K Operations
 */

// @lc code=start
import pq "github.com/emirpasic/gods/v2/queues/priorityqueue"

func maxKelements(nums []int, k int) int64 {
	q := pq.NewWith(func(x, y int) int {
		return y - x
	})
	for _, v := range nums {
		q.Enqueue(v)
	}

	ans := 0
	for i := 0; i < k; i++ {
		v, _ := q.Dequeue()
		ans += v
		q.Enqueue((v + 2) / 3)
	}

	return int64(ans)
}

// @lc code=end
