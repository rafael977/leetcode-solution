package leetcodesolution

/*
 * @lc app=leetcode id=2134 lang=golang
 *
 * [2134] Minimum Swaps to Group All 1's Together II
 */

// @lc code=start
func minSwaps(nums []int) (ans int) {
	n := len(nums)
	ones := 0
	for _, v := range nums {
		ones += v
	}
	ans = ones
	count := 0
	for i := 0; i < n+ones; i++ {
		count += nums[i%n]
		if i < ones-1 {
			continue
		}
		if i >= ones {
			count -= nums[(i-ones)%n]
		}
		ans = min(ans, ones-count)
	}
	return
}

// @lc code=end
