package lettercombinationsofaphonenumber

/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
func combine(a, b []string) (r []string) {
	if len(a) == 0 {
		return b
	}

	for _, sa := range a {
		for _, sb := range b {
			r = append(r, sa+sb)
		}
	}
	return
}

func letterCombinations(digits string) []string {
	m := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	ans := make([]string, 0)
	for _, d := range []byte(digits) {
		ans = combine(ans, m[d])
	}

	return ans
}

// @lc code=end
