package main

import (
	"strings"
)

/*
 * @lc app=leetcode id=71 lang=golang
 *
 * [71] Simplify Path
 */

// @lc code=start
func simplifyPath(path string) string {
	ds := strings.FieldsFunc(path, func(r rune) bool {
		return r == rune('/')
	})

	stack := make([]string, 0, len(ds))
	for _, v := range ds {
		switch v {
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		case ".":
			continue
		default:
			stack = append(stack, v)
		}
	}
	return "/" + strings.Join(stack, "/")
}

// @lc code=end
