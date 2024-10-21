package leetcodesolution

/*
 * @lc app=leetcode id=1701 lang=golang
 *
 * [1701] Average Waiting Time
 */

// @lc code=start
func averageWaitingTime(customers [][]int) float64 {
	time := 0
	total := 0
	for _, v := range customers {
		if v[0] > time {
			time = v[0] + v[1]
		} else {
			time += v[1]
		}
		total += time - v[0]
	}

	return float64(total) / float64(len(customers))
}

// @lc code=end
