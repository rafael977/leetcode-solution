package main

import "sort"

/*
 * @lc app=leetcode id=1337 lang=golang
 *
 * [1337] The K Weakest Rows in a Matrix
 */

// @lc code=start
func kWeakestRows(mat [][]int, k int) (ans []int) {
	numSoldiers := func(r []int) int {
		for i, v := range r {
			if v == 0 {
				return i
			}
		}
		return len(r)
	}

	rows := []struct {
		idx int
		ns  int
	}{}
	for i, v := range mat {
		rows = append(rows, struct {
			idx int
			ns  int
		}{
			idx: i,
			ns:  numSoldiers(v),
		})
	}
	sort.SliceStable(rows, func(i, j int) bool {
		return rows[i].ns < rows[j].ns
	})

	for i := 0; i < k; i++ {
		ans = append(ans, rows[i].idx)
	}

	return
}

// @lc code=end
