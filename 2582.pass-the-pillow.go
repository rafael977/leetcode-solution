package leetcodesolution

/*
 * @lc app=leetcode id=2582 lang=golang
 *
 * [2582] Pass the Pillow
 */

// @lc code=start
func passThePillow(n int, time int) int {
	forward := true
	for time > n-1 {
		time -= n - 1
		forward = !forward
	}
	if forward {
		return 1 + time
	}
	return n - time
}

// @lc code=end
