package main

/*
 * @lc app=leetcode id=532 lang=golang
 *
 * [532] K-diff Pairs in an Array
 */

// @lc code=start
func findPairs(nums []int, k int) (ans int) {
	numMap := make(map[int]bool)
	pair := make(map[int]bool) // record smaller number of a pair

	for _, n := range nums {
		if numMap[n-k] && !pair[n-k] {
			ans++
			pair[n-k] = true
		}
		if numMap[n+k] && !pair[n] {
			ans++
			pair[n] = true
		}

		numMap[n] = true
	}
	return
}

// @lc code=end
