package main

/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 */

// @lc code=start
func subarraySum(nums []int, k int) (ans int) {
	sums := make(map[int]int)
	sums[0] = 1
	sum := 0
	for _, n := range nums {
		sum += n
		ans += sums[sum-k]
		sums[sum]++
	}
	return
}

// @lc code=end
