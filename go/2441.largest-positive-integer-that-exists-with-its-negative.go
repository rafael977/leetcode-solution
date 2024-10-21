package leetcodesolution

/*
 * @lc app=leetcode id=2441 lang=golang
 *
 * [2441] Largest Positive Integer That Exists With Its Negative
 */

// @lc code=start
func findMaxK(nums []int) int {
	neg := make(map[int]struct{})
	for _, v := range nums {
		if v < 0 {
			neg[v] = struct{}{}
		}
	}

	ans := -1
	for _, v := range nums {
		if _, has := neg[-v]; v > 0 && v > ans && has {
			ans = v
		}
	}
	return ans
}

// @lc code=end
