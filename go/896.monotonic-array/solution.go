package monotonicarray

/*
 * @lc app=leetcode id=896 lang=golang
 *
 * [896] Monotonic Array
 */

// @lc code=start
func isMonotonic(nums []int) bool {
	check := func(less func(i, j int) bool) bool {
		for i := 0; i < len(nums)-1; i++ {
			if !less(nums[i], nums[i+1]) {
				return false
			}
		}
		return true
	}

	return check(func(i, j int) bool { return i <= j }) ||
		check(func(i, j int) bool { return i >= j })
}

// @lc code=end
