package shortestpathinbinarymatrix

/*
 * @lc app=leetcode id=1091 lang=golang
 *
 * [1091] Shortest Path in Binary Matrix
 */

// @lc code=start
type index struct {
	x, y int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}

	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
	}
	dist[0][0] = 1
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	visited[0][0] = true

	q := []index{{x: 0, y: 0}}
	for len(q) > 0 {
		idx := q[0]
		q = q[1:]

		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				x, y := idx.x+dx, idx.y+dy

				if (dx == 0 && dy == 0) ||
					x < 0 || y < 0 ||
					x >= n || y >= n {
					continue
				}

				if !visited[x][y] && grid[x][y] == 0 {
					dist[x][y] = dist[idx.x][idx.y] + 1
					visited[x][y] = true
					q = append(q, index{x: x, y: y})
				}
			}
		}
	}

	if dist[n-1][n-1] == 0 {
		return -1
	}
	return dist[n-1][n-1]
}

// @lc code=end
