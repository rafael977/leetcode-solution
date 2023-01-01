package climbingstairs

/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start

var mem map[int]int = make(map[int]int)

func climbStairs(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}

	if c, ok := mem[n]; ok {
		return c
	}
	c := climbStairs(n-1) + climbStairs(n-2)
	mem[n] = c
	return c
}

// @lc code=end
