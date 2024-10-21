package mirrorreflection

/*
 * @lc app=leetcode id=858 lang=golang
 *
 * [858] Mirror Reflection
 */

// @lc code=start
func mirrorReflection(p int, q int) int {
	m := lcm(p, q)
	pt, qt := m/q, m/p

	if pt&1 == 0 {
		return 2
	}
	if qt&1 == 0 {
		return 0
	} else {
		return 1
	}
}

func lcm(p, q int) int {
	m := q
	for m%p != 0 {
		m += q
	}
	return m
}

// @lc code=end
