package main

import "sort"

/*
 * @lc app=leetcode id=1249 lang=golang
 *
 * [1249] Minimum Remove to Make Valid Parentheses
 */

// @lc code=start
func minRemoveToMakeValid(s string) string {
	stack := make([]int, 0, len(s))
	toRemove := make([]int, 0, len(s))
	for i, v := range s {
		if v == '(' {
			stack = append(stack, i)
		} else if v == ')' {
			if len(stack) == 0 {
				toRemove = append(toRemove, i)
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	toRemove = append(toRemove, stack...)
	sort.Slice(toRemove, func(i, j int) bool {
		return toRemove[i] > toRemove[j]
	})

	ss := []rune(s)
	for _, i := range toRemove {
		ss = append(ss[:i], ss[i+1:]...)
	}
	return string(ss)
}

// @lc code=end
