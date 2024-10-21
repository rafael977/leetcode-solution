package minimumnumberofswapstomakethestringbalanced

/*
 * @lc app=leetcode id=1963 lang=golang
 *
 * [1963] Minimum Number of Swaps to Make the String Balanced
 */

// @lc code=start
func minSwaps(s string) int {
	st := make([]rune, 0, len(s))
	for _, v := range s {
		if v == '[' {
			st = append(st, v)
		} else if v == ']' {
			if len(st) > 0 && st[len(st)-1] == '[' {
				st = st[:len(st)-1]
			} else {
				st = append(st, v)
			}
		}
	}

	l := len(st) / 2
	return l/2 + l%2
}

// @lc code=end
