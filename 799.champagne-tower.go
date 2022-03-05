package main

/*
 * @lc app=leetcode id=799 lang=golang
 *
 * [799] Champagne Tower
 */

// @lc code=start
func champagneTower(poured int, query_row int, query_glass int) float64 {
	index := func(i, j int) int {
		if i == 0 {
			return 0
		}
		return (1+i)*i/2 + j
	}

	lc := index(query_row, query_glass)
	pouredLeft := make([]float64, lc+1)
	pouredLeft[0] = float64(poured)

	i, j := 1, 0
	for index(i, j) <= lc {
		q := 0.0
		if j-1 >= 0 {
			l := (pouredLeft[index(i-1, j-1)] - 1) / 2
			if l > 0 {
				q += l
			}
		}
		if j < i {
			r := (pouredLeft[index(i-1, j)] - 1) / 2
			if r > 0 {
				q += r
			}
		}
		pouredLeft[index(i, j)] = q

		if j == i {
			i++
			j = 0
		} else {
			j++
		}
	}

	if pouredLeft[lc] > 1 {
		return 1
	}
	return pouredLeft[lc]
}

// @lc code=end
