package main

/*
 * @lc app=leetcode id=986 lang=golang
 *
 * [986] Interval List Intersections
 */

// @lc code=start
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	li, lj := len(firstList), len(secondList)
	i, j := 0, 0
	ans := make([][]int, 0)
	for i < li && j < lj {
		lo := firstList[i][0]
		if secondList[j][0] > firstList[i][0] {
			lo = secondList[j][0]
		}
		hi := firstList[i][1]
		if secondList[j][1] < firstList[i][1] {
			hi = secondList[j][1]
		}
		if lo <= hi {
			ans = append(ans, []int{lo, hi})
		}
		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}

	return ans
}

// @lc code=end
