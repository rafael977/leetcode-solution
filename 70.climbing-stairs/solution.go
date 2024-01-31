package climbingstairs

/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start

// var mem map[int]int = make(map[int]int)

func climbStairs(n int) (ans int) {
	// if n == 0 {
	// 	return 1
	// }
	// if n < 0 {
	// 	return 0
	// }

	// if c, ok := mem[n]; ok {
	// 	return c
	// }
	// c := climbStairs(n-1) + climbStairs(n-2)
	// mem[n] = c
	// return c
	i1, i2 := 1, 1
	ans = 1
	for i := 2; i <= n; i++ {
		ans = i1 + i2
		i2 = i1
		i1 = ans
	}
	return
}

// @lc code=end
