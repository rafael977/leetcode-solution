package main

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode id=402 lang=golang
 *
 * [402] Remove K Digits
 */

// @lc code=start
func removeKdigits(num string, k int) string {
	stack := make([]int, 0, len(num))
	for _, r := range num {
		n, _ := strconv.Atoi(string(r))
		for len(stack) > 0 && k > 0 && stack[len(stack)-1] > n {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, n)
	}

	s := ""
	for i := 0; i < len(stack)-k; i++ {
		if s == "" && stack[i] == 0 {
			continue
		}
		s += fmt.Sprintf("%d", stack[i])
	}
	if s == "" {
		return "0"
	}
	return s
}

// @lc code=end
