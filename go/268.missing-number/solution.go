package missingnumber

/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */

// @lc code=start
func missingNumber(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return total - sum
}

// @lc code=end
