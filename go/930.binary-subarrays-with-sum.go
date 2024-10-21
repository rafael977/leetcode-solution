package leetcodesolution

/*
 * @lc app=leetcode id=930 lang=golang
 *
 * [930] Binary Subarrays With Sum
 */

// @lc code=start
func numSubarraysWithSum(nums []int, goal int) (ans int) {
	cnt := make(map[int]int)
	sum := 0
	for _, n := range nums {
		cnt[sum]++
		sum += n
		ans += cnt[sum-goal]
	}
	return
}

// @lc code=end
