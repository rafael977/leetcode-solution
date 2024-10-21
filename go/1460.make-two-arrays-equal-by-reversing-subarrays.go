package leetcodesolution

/*
 * @lc app=leetcode id=1460 lang=golang
 *
 * [1460] Make Two Arrays Equal by Reversing Subarrays
 */

// @lc code=start
func canBeEqual(target []int, arr []int) bool {
	nc := make(map[int]int)
	for _, v := range target {
		nc[v]++
	}

	for _, v := range arr {
		nc[v]--
	}

	for _, v := range nc {
		if v != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
