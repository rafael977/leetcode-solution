package wigglesubsequence

/*
 * @lc app=leetcode id=376 lang=golang
 *
 * [376] Wiggle Subsequence
 */

// @lc code=start
func wiggleMaxLength(nums []int) (ans int) {
	n := len(nums)
	if n < 2 {
		return n
	}
	ans = 1
	prevDiff := nums[1] - nums[0]
	if prevDiff != 0 {
		ans = 2
	}
	for i := 2; i < n; i++ {
		diff := nums[i] - nums[i-1]
		if diff > 0 && prevDiff <= 0 || diff < 0 && prevDiff >= 0 {
			ans++
			prevDiff = diff
		}
	}
	return
}

// @lc code=end
