package leetcodesolution

/*
 * @lc app=leetcode id=1550 lang=golang
 *
 * [1550] Three Consecutive Odds
 */

// @lc code=start
func threeConsecutiveOdds(arr []int) bool {
	for i := 2; i < len(arr); i++ {
		if arr[i-2]&1 == 1 && arr[i-1]&1 == 1 && arr[i]&1 == 1 {
			return true
		}
	}
	return false
}

// @lc code=end
