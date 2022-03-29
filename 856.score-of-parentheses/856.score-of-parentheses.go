package main

/*
 * @lc app=leetcode id=856 lang=golang
 *
 * [856] Score of Parentheses
 */

// @lc code=start
func scoreOfParentheses(s string) int {
	score := 0

	i := 0
	stack := 0
	for j, v := range s {
		if v == '(' {
			stack++
		} else {
			stack--
		}

		if stack == 0 {
			if j-i == 1 {
				score++
			} else {
				score += 2 * scoreOfParentheses(s[i+1:j])
			}
			i = j + 1
		}
	}
	return score
}

// @lc code=end
