package leetcodesolution

/*
 * @lc app=leetcode id=2958 lang=golang
 *
 * [2958] Length of Longest Subarray With at Most K Frequency
 */

// @lc code=start
func maxSubarrayLength(nums []int, k int) (ans int) {
	cnt := make(map[int]int)
	for i, j := 0, 0; j < len(nums); {
		cnt[nums[j]]++
		for i < j && cnt[nums[j]] > k {
			cnt[nums[i]]--
			i++
		}
		ans = max(ans, j-i+1)
		j++
	}
	return
}

// @lc code=end
