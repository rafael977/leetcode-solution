package fractionadditionandsubtraction

import (
	"fmt"
	"unicode"
)

/*
 * @lc app=leetcode id=592 lang=golang
 *
 * [592] Fraction Addition and Subtraction
 */

// @lc code=start
func fractionAddition(expression string) (ans string) {
	x, y := 0, 1
	for i, n := 0, len(expression); i < n; {
		x1, sign := 0, 1
		if expression[i] == '-' || expression[i] == '+' {
			if expression[i] == '-' {
				sign = -1
			}
			i++
		}
		for i < n && unicode.IsDigit(rune(expression[i])) {
			x1 = x1*10 + int(expression[i]-'0')
			i++
		}
		x1 = sign * x1
		i++

		y1 := 0
		for i < n && unicode.IsDigit(rune(expression[i])) {
			y1 = y1*10 + int(expression[i]-'0')
			i++
		}

		x = x*y1 + x1*y
		y *= y1
	}

	px := x
	if x < 0 {
		px = -x
	}
	g := gcd(px, y)
	return fmt.Sprintf("%d/%d", x/g, y/g)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// @lc code=end
