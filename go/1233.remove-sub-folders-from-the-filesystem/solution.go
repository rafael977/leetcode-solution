// Created by Rafael Shen at 2024/10/25 11:55
// leetgo: 1.4.9
// https://leetcode.com/problems/remove-sub-folders-from-the-filesystem/

package removesubfoldersfromthefilesystem

import (
	"sort"
	"strings"
)

// @lc code=begin

func removeSubfolders(folder []string) (ans []string) {
	sort.Strings(folder)
	ans = append(ans, folder[0])
	for _, f := range folder[1:] {
		last := ans[len(ans)-1]
		if !strings.HasPrefix(f, last) || f[len(last)] != '/' {
			ans = append(ans, f)
		}
	}
	return
}

// @lc code=end
