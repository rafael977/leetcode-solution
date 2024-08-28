package countsubislands

/*
 * @lc app=leetcode id=1905 lang=golang
 *
 * [1905] Count Sub Islands
 */

// @lc code=start
func countSubIslands(grid1 [][]int, grid2 [][]int) (ans int) {
	m, n := len(grid1), len(grid1[0])
	bfs := func(x, y int) int {
		q := [][2]int{{x, y}}
		in := 1
		if grid1[x][y] == 0 {
			in = 0
		}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			if grid1[p[0]][p[1]] == 0 {
				in = 0
			}
			for _, d := range [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} {
				xx, yy := p[0]+d[0], p[1]+d[1]
				if xx >= 0 && xx < m && yy >= 0 && yy < n && grid2[xx][yy] == 1 {
					grid2[xx][yy] = 0
					q = append(q, [2]int{xx, yy})
				}
			}
		}
		return in
	}

	for i := range grid2 {
		for j := range grid2[i] {
			if grid2[i][j] == 1 {
				ans += bfs(i, j)
			}
		}
	}

	return
}

// @lc code=end
