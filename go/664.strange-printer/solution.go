package strangeprinter

/*
 * @lc app=leetcode id=664 lang=golang
 *
 * [664] Strange Printer
 */

// @lc code=start
func strangePrinter(s string) int {
	ss := make([]byte, 0, len(s))
	c := s[0]
	ss = append(ss, c)
	for _, v := range []byte(s) {
		if v != c {
			ss = append(ss, v)
			c = v
		}
	}

	dp := make([][]int, len(ss))
	for i := range dp {
		dp[i] = make([]int, len(ss))
		for j := range dp[i] {
			dp[i][j] = len(s)
		}
		dp[i][i] = 1
	}

	for i := len(ss) - 1; i >= 0; i-- {
		for j := i + 1; j < len(ss); j++ {
			if ss[i] == ss[j] {
				dp[i][j] = dp[i][j-1]
				continue
			}
			for k := i; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j])
			}
		}
	}

	return dp[0][len(ss)-1]
}

// @lc code=end
