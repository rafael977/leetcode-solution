package topkfrequentwords

import "sort"

/*
 * @lc app=leetcode id=692 lang=golang
 *
 * [692] Top K Frequent Words
 */

// @lc code=start
type wordCount struct {
	word  string
	count int
}

func topKFrequent(words []string, k int) []string {
	m := make(map[string]*wordCount)
	for _, w := range words {
		if m[w] == nil {
			m[w] = &wordCount{word: w, count: 1}
		} else {
			m[w].count++
		}
	}

	wcs := make([]*wordCount, len(m))
	i := 0
	for _, wc := range m {
		wcs[i] = wc
		i++
	}

	sort.Slice(wcs, func(i, j int) bool {
		if wcs[i].count == wcs[j].count {
			return wcs[i].word < wcs[j].word
		}
		return wcs[i].count > wcs[j].count
	})
	ans := make([]string, k)
	for i := 0; i < k; i++ {
		ans[i] = wcs[i].word
	}
	return ans
}

// @lc code=end
