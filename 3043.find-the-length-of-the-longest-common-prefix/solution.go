package findthelengthofthelongestcommonprefix

import "fmt"

/*
 * @lc app=leetcode id=3043 lang=golang
 *
 * [3043] Find the Length of the Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(arr1 []int, arr2 []int) (ans int) {
	pm := make(map[int]bool)
	for _, v := range arr1 {
		for v > 0 {
			pm[v] = true
			v /= 10
		}
	}
	for _, v := range arr2 {
		for v > 0 {
			if pm[v] {
				ans = max(ans, len(fmt.Sprintf("%d", v)))
				break
			}
			v /= 10
		}
	}
	return
}

// @lc code=end
