package leetcodesolution

/*
 * @lc app=leetcode id=2812 lang=golang
 *
 * [2812] Find the Safest Path in a Grid
 */

// @lc code=start

func maximumSafenessFactor(grid [][]int) (ans int) {
	n := len(grid)
	type pair struct{ x, y int }
	q := []pair{}
	dis := make([][]int, n)
	for i, row := range grid {
		dis[i] = make([]int, n)
		for j, x := range row {
			if x > 0 {
				q = append(q, pair{i, j})
			} else {
				dis[i][j] = -1
			}
		}
	}
	dir4 := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	groups := [][]pair{q}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, p := range tmp {
			for _, d := range dir4 {
				x, y := p.x+d.x, p.y+d.y
				if 0 <= x && x < n && 0 <= y && y < n && dis[x][y] < 0 {
					q = append(q, pair{x, y})
					dis[x][y] = len(groups)
				}
			}
		}

		groups = append(groups, q)
	}

	fa := make([]int, n*n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	for ans := len(groups) - 2; ans > 0; ans-- {
		for _, p := range groups[ans] {
			i, j := p.x, p.y
			for _, d := range dir4 {
				x, y := p.x+d.x, p.y+d.y
				if 0 <= x && x < n && 0 <= y && y < n && dis[x][y] >= dis[i][j] {
					fa[find(x*n+y)] = find(i*n + j)
				}
			}
		}

		if find(0) == find(n*n-1) {
			return ans
		}
	}

	return 0
}

// @lc code=end
