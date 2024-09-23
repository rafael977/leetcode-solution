package extracharactersinastring

/*
 * @lc app=leetcode id=2707 lang=golang
 *
 * [2707] Extra Characters in a String
 */

// @lc code=start
func minExtraChar(s string, dictionary []string) int {
	dict := make(map[string]bool)
	for _, v := range dictionary {
		dict[v] = true
	}

	memo := make([]int, len(s))
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}

		if memo[i] != -1 {
			return memo[i]
		}

		r := dfs(i-1) + 1
		for j := 0; j <= i; j++ {
			if dict[s[j:i+1]] {
				r = min(r, dfs(j-1))
			}
		}
		memo[i] = r
		return r
	}

	return dfs(len(s) - 1)
}

// @lc code=end
