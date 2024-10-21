package sortcharactersbyfrequency

import (
	"sort"
	"strings"
)

/*
 * @lc app=leetcode id=451 lang=golang
 *
 * [451] Sort Characters By Frequency
 */

// @lc code=start
func frequencySort(s string) string {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}

	type CharFreq struct {
		Char rune
		Freq int
	}
	cf := make([]CharFreq, 0)
	for k, v := range m {
		cf = append(cf, CharFreq{Char: k, Freq: v})
	}

	sort.Slice(cf, func(i, j int) bool {
		if cf[i].Freq == cf[j].Freq {
			return cf[i].Char < cf[j].Char
		}
		return cf[i].Freq > cf[j].Freq
	})
	sb := strings.Builder{}
	for _, v := range cf {
		for i := 0; i < v.Freq; i++ {
			sb.WriteRune(v.Char)
		}
	}
	return sb.String()
}

// @lc code=end
