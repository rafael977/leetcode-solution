package main

/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) (ans [][]int) {
	for mask := 0; mask < 1<<len(nums); mask++ {
		set := []int{}
		for i, n := range nums {
			if mask>>i&1 > 0 {
				set = append(set, n)
			}
		}
		ans = append(ans, set)
	}
	return
}

// @lc code=end
