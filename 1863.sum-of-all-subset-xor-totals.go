package leetcodesolution

/*
 * @lc app=leetcode id=1863 lang=golang
 *
 * [1863] Sum of All Subset XOR Totals
 */

// @lc code=start
func subsetXORSum(nums []int) (ans int) {
	var dfs func(sum, idx int)
	dfs = func(sum, idx int) {
		if idx == len(nums) {
			ans += sum
			return
		}

		dfs(sum, idx+1)
		dfs(sum^nums[idx], idx+1)
	}

	dfs(0, 0)
	return
}

// @lc code=end
