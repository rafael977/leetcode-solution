package findduplicatefileinsystem

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=609 lang=golang
 *
 * [609] Find Duplicate File in System
 */

// @lc code=start
func findDuplicate(paths []string) [][]string {
	dic := make(map[string][]string)
	for _, path := range paths {
		parseDir(path, dic)
	}

	ans := make([][]string, 0)
	for _, d := range dic {
		if len(d) > 1 {
			ans = append(ans, d)
		}
	}
	return ans
}

func parseDir(path string, dic map[string][]string) {
	parts := strings.Fields(path)
	dir := parts[0]
	for _, f := range parts[1:] {
		fs := strings.Split(f, "(")
		content := fs[1][:len(fs[1])-1]
		if _, ok := dic[content]; !ok {
			dic[content] = make([]string, 0)
		}
		dic[content] = append(dic[content], fmt.Sprintf("%s/%s", dir, fs[0]))
	}
}

// @lc code=end
