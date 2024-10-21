package minimumdeletionstomakecharacterfrequenciesunique

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=1647 lang=golang
 *
 * [1647] Minimum Deletions to Make Character Frequencies Unique
 */

// @lc code=start
func minDeletions(s string) (ans int) {
	freq := make([]int, 26)
	for _, c := range s {
		freq[c-'a']++
	}

	sort.Slice(freq, func(i, j int) bool {
		return freq[i] > freq[j]
	})

	avil := math.MaxInt64
	for _, f := range freq {
		if f == 0 {
			break
		}

		if avil == 0 {
			ans += f
			continue
		}

		if f > avil {
			ans += f - avil
			avil--
		} else {
			avil = f - 1
		}
	}
	return
}

// @lc code=end
