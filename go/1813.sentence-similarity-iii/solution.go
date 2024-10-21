package sentencesimilarityiii

import "strings"

/*
 * @lc app=leetcode id=1813 lang=golang
 *
 * [1813] Sentence Similarity III
 */

// @lc code=start
func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	words1, words2 := strings.Fields(sentence1), strings.Fields(sentence2)
	n, m := len(words1), len(words2)
	i, j := 0, 0
	for i < n && i < m && words1[i] == words2[i] {
		i++
	}
	for j < n-i && j < m-i && words1[n-1-j] == words2[m-1-j] {
		j++
	}

	return i+j == min(m, n)
}

// @lc code=end
