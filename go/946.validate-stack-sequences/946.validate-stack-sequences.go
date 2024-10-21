package main

/*
 * @lc app=leetcode id=946 lang=golang
 *
 * [946] Validate Stack Sequences
 */

// @lc code=start
func validateStackSequences(pushed []int, popped []int) bool {
	n := len(pushed)
	stack := make([]int, 0, n)
	j := 0
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && j < n && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return j == n
}

// @lc code=end
