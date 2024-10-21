package leetcodesolution

/*
 * @lc app=leetcode id=1518 lang=golang
 *
 * [1518] Water Bottles
 */

// @lc code=start
func numWaterBottles(numBottles int, numExchange int) int {
	ans := numBottles
	for numBottles >= numExchange {
		exc := numBottles / numExchange
		numBottles = exc + numBottles%numExchange
		ans += exc
	}

	return ans
}

// @lc code=end
