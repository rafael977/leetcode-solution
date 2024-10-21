package leetcodesolution

import "strings"

/*
 * @lc app=leetcode id=140 lang=golang
 *
 * [140] Word Break II
 */

// @lc code=start
func wordBreak(s string, wordDict []string) (ans []string) {
	wd := make(map[string]struct{})
	for _, v := range wordDict {
		wd[v] = struct{}{}
	}

	var backtrack func(i int, words []string)
	backtrack = func(i int, words []string) {
		if i == len(s) {
			ans = append(ans, strings.Join(words, " "))
			return
		}
		for j := i + 1; j <= len(s); j++ {
			word := s[i:j]
			if _, in := wd[word]; in {
				newWords := append([]string{}, words...)
				newWords = append(newWords, word)
				backtrack(j, newWords)
			}
		}
	}

	backtrack(0, []string{})
	return
}

// @lc code=end
