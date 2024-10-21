package leetcodesolution

/*
 * @lc app=leetcode id=959 lang=golang
 *
 * [959] Regions Cut By Slashes
 */

// @lc code=start
func regionsBySlashes(grid []string) int {
	n := len(grid)
	parent := make([]int, 4*n*n)
	count := 4 * n * n
	for i := range parent {
		parent[i] = i
	}

	find := func(x int) int {
		for x != parent[x] {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}
	union := func(x, y int) {
		px, py := find(x), find(y)
		if px != py {
			parent[py] = px
			count--
		}
	}

	for i, row := range grid {
		for j, c := range row {
			idx := 4 * (i*n + j)
			if c == '/' {
				union(idx, idx+3)
				union(idx+1, idx+2)
			} else if c == '\\' {
				union(idx, idx+1)
				union(idx+2, idx+3)
			} else {
				union(idx, idx+1)
				union(idx+1, idx+2)
				union(idx+2, idx+3)
			}

			if j+1 < n {
				union(idx+1, 4*(i*n+j+1)+3)
			}
			if i+1 < n {
				union(idx+2, 4*((i+1)*n+j))
			}
		}
	}
	return count
}

// @lc code=end
