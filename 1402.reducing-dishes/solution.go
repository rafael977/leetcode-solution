package reducingdishes

import "sort"

/*
 * @lc app=leetcode id=1402 lang=golang
 *
 * [1402] Reducing Dishes
 */

// @lc code=start
func maxSatisfaction(satisfaction []int) (ans int) {
	sort.Ints(satisfaction)
	sum, score := 0, 0
	for i := len(satisfaction) - 1; i >= 0; i-- {
		score += sum + satisfaction[i]
		sum += satisfaction[i]
		ans = max(ans, score)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
