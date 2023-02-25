package subarraysumsdivisiblebyk

/*
 * @lc app=leetcode id=974 lang=golang
 *
 * [974] Subarray Sums Divisible by K
 */

// @lc code=start
func subarraysDivByK(nums []int, k int) (ans int) {
	rc := map[int]int{0: 1}
	sum := 0
	for _, v := range nums {
		sum += v
		r := (sum%k + k) % k
		ans += rc[r]
		rc[r]++
	}
	return
}

// @lc code=end
