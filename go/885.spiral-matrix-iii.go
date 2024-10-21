package leetcodesolution

/*
 * @lc app=leetcode id=885 lang=golang
 *
 * [885] Spiral Matrix III
 */

// @lc code=start
func spiralMatrixIII(rows int, cols int, rStart int, cStart int) (ans [][]int) {
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	d, s, steps := 0, 0, 1
	for len(ans) < rows*cols {
		if rStart >= 0 && rStart < rows && cStart >= 0 && cStart < cols {
			ans = append(ans, []int{rStart, cStart})
		}
		rStart += dirs[d][0]
		cStart += dirs[d][1]
		s++
		if s == steps {
			s = 0
			d++
			if d == 2 || d == 4 {
				d = d % 4
				steps++
			}
		}
	}
	return
}

// @lc code=end
