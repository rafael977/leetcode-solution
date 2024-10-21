package coinchange

/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	dp := make(map[int]int)
	var change func(a int) int
	change = func(a int) int {
		if a < 0 {
			return -1
		}
		if a == 0 {
			return 0
		}
		if cc, ok := dp[a]; ok {
			return cc
		}

		min := -1
		for _, c := range coins {
			cc := change(a - c)
			if cc == -1 {
				continue
			}
			if min == -1 || min > cc+1 {
				min = cc + 1
			}
		}
		dp[a] = min
		return min
	}

	return change(amount)
}

// @lc code=end
