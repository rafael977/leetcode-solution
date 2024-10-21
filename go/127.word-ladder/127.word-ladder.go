package main

import "math"

/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	q := []string{beginWord}
	ans := 1
	for len(wordList) > 0 && len(q) > 0 {
		ans++
		tmpQ := make([]string, 0, len(wordList))
		for _, s := range q {
			remove := make([]int, 0, len(wordList))
			for i, w := range wordList {
				if diff(s, w) == 1 {
					if w == endWord {
						return ans
					}
					tmpQ = append(tmpQ, w)
					remove = append(remove, i)
				}
			}
			for i := len(remove) - 1; i >= 0; i-- {
				wordList = append(wordList[:remove[i]], wordList[remove[i]+1:]...)
			}
		}
		q = tmpQ
	}
	return 0
}

func diff(a, b string) (r int) {
	if len(a) != len(b) {
		return math.MaxInt64
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			r++
		}
	}
	return
}

// @lc code=end
