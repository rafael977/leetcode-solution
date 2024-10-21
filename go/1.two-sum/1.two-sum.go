package main

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	idxMap := make(map[int]int)
	ans := make([]int, 2)
	for i, v := range nums {
		if idx, ok := idxMap[v]; ok {
			ans = []int{idx, i}
		}

		idxMap[target-v] = i
	}

	return ans
}

// @lc code=end
