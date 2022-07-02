package maximumareaofapieceofcakeafterhorizontalandverticalcuts

import "sort"

/*
 * @lc app=leetcode id=1465 lang=golang
 *
 * [1465] Maximum Area of a Piece of Cake After Horizontal and Vertical Cuts
 */

// @lc code=start
func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	const modulo = 1_000_000_000 + 7

	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)

	hh, ww := 0, 0
	for i := 0; i <= len(horizontalCuts); i++ {
		hi := 0
		if i == 0 {
			hi = horizontalCuts[i]
		} else if i == len(horizontalCuts) {
			hi = h - horizontalCuts[i-1]
		} else {
			hi = horizontalCuts[i] - horizontalCuts[i-1]
		}
		hh = max(hh, hi)
	}

	for j := 0; j <= len(verticalCuts); j++ {
		wi := 0
		if j == 0 {
			wi = verticalCuts[j]
		} else if j == len(verticalCuts) {
			wi = w - verticalCuts[j-1]
		} else {
			wi = verticalCuts[j] - verticalCuts[j-1]
		}

		ww = max(ww, wi)
	}

	return (hh * ww) % modulo
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
