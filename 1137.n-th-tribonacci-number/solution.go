package nthtribonaccinumber

/*
 * @lc app=leetcode id=1137 lang=golang
 *
 * [1137] N-th Tribonacci Number
 */

// @lc code=start
func tribonacci(n int) int {
	m := map[int]int{
		0: 0,
		1: 1,
		2: 1,
	}
	var calc func(n int) int
	calc = func(n int) int {
		if _, ok := m[n]; ok {
			return m[n]
		}
		r := calc(n-1) + calc(n-2) + calc(n-3)
		m[n] = r
		return r
	}

	return calc(n)
}

// @lc code=end
