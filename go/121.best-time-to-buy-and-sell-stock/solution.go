package besttimetobuyandsellstock

/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) (ans int) {
	min := prices[0]
	for _, v := range prices[1:] {
		if v-min > ans {
			ans = v - min
		}
		if v < min {
			min = v
		}
	}
	return
}

// @lc code=end
