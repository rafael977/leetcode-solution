package bagoftokens

import "sort"

/*
 * @lc app=leetcode id=948 lang=golang
 *
 * [948] Bag of Tokens
 */

// @lc code=start
func bagOfTokensScore(tokens []int, power int) (ans int) {
	sort.Ints(tokens)
	i, j := 0, len(tokens)-1
	score := 0
	for i <= j {
		if power >= tokens[i] {
			score++
			power -= tokens[i]
			i++
		} else if score > 0 {
			score--
			power += tokens[j]
			j--
		} else {
			break
		}

		ans = max(ans, score)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
