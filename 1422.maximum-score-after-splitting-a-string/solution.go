package maximumscoreaftersplittingastring

/*
 * @lc app=leetcode id=1422 lang=golang
 *
 * [1422] Maximum Score After Splitting a String
 */

// @lc code=start
func maxScore(s string) (ans int) {
	score := 0
	for _, c := range s {
		if c == '1' {
			score++
		}
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			score++
		} else if s[i] == '1' {
			score--
		}

		if ans < score {
			ans = score
		}
	}
	return
}

// @lc code=end
