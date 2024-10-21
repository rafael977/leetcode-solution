package leetcodesolution

import "math/bits"

/*
 * @lc app=leetcode id=2997 lang=golang
 *
 * [2997] Minimum Number of Operations to Make Array XOR Equal to K
 */

// @lc code=start
func minOperations(nums []int, k int) int {
	for _, v := range nums {
		k ^= v
	}
	return bits.OnesCount(uint(k))
}

// @lc code=end
