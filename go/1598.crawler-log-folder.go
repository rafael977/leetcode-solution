package leetcodesolution

/*
 * @lc app=leetcode id=1598 lang=golang
 *
 * [1598] Crawler Log Folder
 */

// @lc code=start
func minOperations1598(logs []string) int {
	lvl := 0
	for _, l := range logs {
		if l == "../" {
			if lvl > 0 {
				lvl--
			}
		} else if l == "./" {
			lvl += 0
		} else {
			lvl++
		}
	}

	return lvl
}

// @lc code=end
