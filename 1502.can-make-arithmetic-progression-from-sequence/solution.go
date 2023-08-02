package canmakearithmeticprogressionfromsequence

import "sort"

/*
 * @lc app=leetcode id=1502 lang=golang
 *
 * [1502] Can Make Arithmetic Progression From Sequence
 */

// @lc code=start
func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}

// @lc code=end
