package leetcodesolution

/*
 * @lc app=leetcode id=2597 lang=golang
 *
 * [2597] The Number of Beautiful Subsets
 */

// @lc code=start
func beautifulSubsets(nums []int, k int) (ans int) {
	isBeautiful := func(ss []int, x int) bool {
		for _, v := range ss {
			if v-x == k || x-v == k {
				return false
			}
		}
		return true
	}

	ss := make([]int, 0)
	var backtrack func(i int)
	backtrack = func(i int) {
		if i == len(nums) {
			if len(ss) > 0 {
				ans++
			}
			return
		}
		if isBeautiful(ss, nums[i]) {
			ss = append(ss, nums[i])
			backtrack(i + 1)
			ss = ss[:len(ss)-1]
		}
		backtrack(i + 1)
	}
	backtrack(0)
	return
}

// @lc code=end
