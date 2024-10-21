package validparentheses

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
			continue
		}

		if len(stack) == 0 ||
			(v == ')' && stack[len(stack)-1] != '(') ||
			(v == '}' && stack[len(stack)-1] != '{') ||
			(v == ']' && stack[len(stack)-1] != '[') {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

// @lc code=end
