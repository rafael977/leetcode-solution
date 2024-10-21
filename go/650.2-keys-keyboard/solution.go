package keyskeyboard

/*
 * @lc app=leetcode id=650 lang=golang
 *
 * [650] 2 Keys Keyboard
 */

// @lc code=start
var dp map[int]int = map[int]int{1: 0}

func minSteps(n int) int {
	if v, has := dp[n]; has {
		return v
	}

	v := n
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			v = min(v, minSteps(i)+n/i)
		}
	}
	dp[n] = v
	return v
}

// @lc code=end
