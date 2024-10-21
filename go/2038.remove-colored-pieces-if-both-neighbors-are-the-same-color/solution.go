package removecoloredpiecesifbothneighborsarethesamecolor

import "strings"

/*
 * @lc app=leetcode id=2038 lang=golang
 *
 * [2038] Remove Colored Pieces if Both Neighbors are the Same Color
 */

// @lc code=start
func winnerOfGame(colors string) bool {
	as := strings.FieldsFunc(colors, func(r rune) bool { return r == 'B' })
	bs := strings.FieldsFunc(colors, func(r rune) bool { return r == 'A' })

	am, bm := 0, 0
	for _, v := range as {
		am += max(len(v)-2, 0)
	}
	for _, v := range bs {
		bm += max(len(v)-2, 0)
	}
	return am > bm
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
