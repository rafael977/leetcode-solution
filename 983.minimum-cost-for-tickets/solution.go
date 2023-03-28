package minimumcostfortickets

/*
 * @lc app=leetcode id=983 lang=golang
 *
 * [983] Minimum Cost For Tickets
 */

// @lc code=start
func mincostTickets(days []int, costs []int) int {
	lastDay := days[len(days)-1]
	dc := make([]int, lastDay+1)
	dc[0] = 0
	i := 0
	for j := 1; j <= lastDay; j++ {
		if j < days[i] {
			dc[j] = dc[j-1]
			continue
		}

		cost1 := costs[0]
		if j-1 >= 0 {
			cost1 += dc[j-1]
		}
		cost7 := costs[1]
		if j-7 >= 0 {
			cost7 += dc[j-7]
		}
		cost30 := costs[2]
		if j-30 >= 0 {
			cost30 += dc[j-30]
		}
		dc[j] = min(cost1, min(cost7, cost30))
		i++
	}
	return dc[lastDay]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
