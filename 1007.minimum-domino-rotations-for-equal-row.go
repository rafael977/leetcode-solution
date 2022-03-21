package main

/*
 * @lc app=leetcode id=1007 lang=golang
 *
 * [1007] Minimum Domino Rotations For Equal Row
 */

// @lc code=start
func minDominoRotations(tops []int, bottoms []int) int {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	n := len(tops)

	ra := 0
	ea := 0
	for i := 1; i < n; i++ {
		if tops[0] != tops[i] && tops[0] != bottoms[i] {
			ra = -1
			break
		}
		if tops[0] == tops[i] && tops[0] == bottoms[i] {
			ea++
		} else if tops[0] == bottoms[i] {
			ra++
		}
	}

	rb := 0
	eb := 0
	for i := 1; i < n; i++ {
		if bottoms[0] != tops[i] && bottoms[0] != bottoms[i] {
			rb = -1
			break
		}
		if bottoms[0] == tops[i] && bottoms[0] == bottoms[i] {
			eb++
		} else if bottoms[0] == tops[i] {
			rb++
		}
	}

	if ra == -1 && rb == -1 {
		return -1
	} else if ra == -1 {
		return min(rb, n-rb-eb)
	} else if rb == -1 {
		return min(ra, n-ra-ea)
	} else {
		return min(min(ra, n-ra-ea), min(rb, n-rb-eb))
	}
}

// @lc code=end
