package evaluatereversepolishnotation

import "strconv"

/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */

// @lc code=start
func evalRPN(tokens []string) int {
	exp := []int{}
	for _, t := range tokens {
		switch t {
		case "+":
			a, b := exp[len(exp)-2], exp[len(exp)-1]
			exp[len(exp)-2] = a + b
			exp = exp[:len(exp)-1]
		case "-":
			a, b := exp[len(exp)-2], exp[len(exp)-1]
			exp[len(exp)-2] = a - b
			exp = exp[:len(exp)-1]
		case "*":
			a, b := exp[len(exp)-2], exp[len(exp)-1]
			exp[len(exp)-2] = a * b
			exp = exp[:len(exp)-1]
		case "/":
			a, b := exp[len(exp)-2], exp[len(exp)-1]
			exp[len(exp)-2] = a / b
			exp = exp[:len(exp)-1]
		default:
			v, _ := strconv.Atoi(t)
			exp = append(exp, v)
		}
	}

	return exp[0]
}

// @lc code=end
