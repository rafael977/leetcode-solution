package main

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=165 lang=golang
 *
 * [165] Compare Version Numbers
 */

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	v1, v2 := strings.Split(version1, "."), strings.Split(version2, ".")
	l := len(v1)
	if l > len(v2) {
		l = len(v2)
	}

	for i := 0; i < l; i++ {
		r1, _ := strconv.Atoi(v1[i])
		r2, _ := strconv.Atoi(v2[i])
		if r1 < r2 {
			return -1
		} else if r1 > r2 {
			return 1
		}
	}
	for i := l; i < len(v1); i++ {
		if r, _ := strconv.Atoi(v1[i]); r != 0 {
			return 1
		}
	}
	for i := l; i < len(v2); i++ {
		if r, _ := strconv.Atoi(v2[i]); r != 0 {
			return -1
		}
	}

	return 0
}

// @lc code=end
