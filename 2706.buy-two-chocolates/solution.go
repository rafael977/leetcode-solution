package buytwochocolates

import "sort"

/*
 * @lc app=leetcode id=2706 lang=golang
 *
 * [2706] Buy Two Chocolates
 */

// @lc code=start
func buyChoco(prices []int, money int) int {
	sort.Ints(prices)
	sum := prices[0] + prices[1]
	if sum <= money {
		return money - sum
	}
	return money
}

// @lc code=end
