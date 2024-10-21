package perfectsquares

import "math"

/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */

// @lc code=start
var dp map[int]int = make(map[int]int)

func numSquares(n int) int {
	if n == 0 {
		return 0
	}
	if v, ok := dp[n]; ok {
		return v
	}

	ans := math.MaxInt
	rt := int(math.Sqrt(float64(n)))
	for i := rt; i > 0; i-- {
		ans = min(ans, numSquares(n-i*i)+1)
	}
	dp[n] = ans
	return ans
}

// @lc code=end
