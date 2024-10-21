package shortencodingofwords

import (
	"sort"
)

/*
 * @lc app=leetcode id=820 lang=golang
 *
 * [820] Short Encoding of Words
 */

// @lc code=start
func minimumLengthEncoding(words []string) (ans int) {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	ws := make([]string, 0, len(words))
	for _, w := range words {
		isSuffix := false
		for _, ww := range ws {
			if suffix(ww, w) {
				isSuffix = true
				break
			}
		}
		if !isSuffix {
			ws = append(ws, w)
		}
	}

	for _, w := range ws {
		ans += len(w) + 1
	}
	return
}

func suffix(a, b string) bool {
	na, nb := len(a), len(b)
	if na < nb {
		a, b = b, a
	}
	for i := range b {
		if a[na-1-i] != b[nb-1-i] {
			return false
		}
	}
	return true
}

// @lc code=end
