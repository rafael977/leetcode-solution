package findsmallestlettergreaterthantarget

import (
	"sort"
)

/*
 * @lc app=leetcode id=744 lang=golang
 *
 * [744] Find Smallest Letter Greater Than Target
 */

// @lc code=start
func nextGreatestLetter(letters []byte, target byte) byte {
	i := sort.Search(len(letters), func(i int) bool {
		return letters[i] > target
	})

	if i == len(letters) {
		i = 0
	}
	return letters[i]
}

// @lc code=end
