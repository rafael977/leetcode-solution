package leetcodesolution

/*
 * @lc app=leetcode id=713 lang=golang
 *
 * [713] Subarray Product Less Than K
 */

// @lc code=start
func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	for j := range nums {
		i := j
		p := 1
		for i >= 0 {
			p *= nums[i]
			if p >= k {
				break
			}
			i--
		}

		ans += (j - i)
	}
	return
}

// @lc code=end
