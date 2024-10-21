package dividetwointegers

/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 */

// @lc code=start
func divide(dividend int, divisor int) (quotient int) {
	sign := 1
	if dividend < 0 {
		dividend = -dividend
		sign = -sign
	}
	if divisor < 0 {
		divisor = -divisor
		sign = -sign
	}
	for dividend >= divisor {
		quotient++
		dividend -= divisor
	}

	if sign == -1 {
		quotient = -quotient
	}
	if quotient > 1<<31-1 {
		return 1<<31 - 1
	}
	if quotient < -(1 << 31) {
		return -(1 << 31)
	}

	return
}

// @lc code=end
