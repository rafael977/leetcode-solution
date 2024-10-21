package networkdelaytime

/*
 * @lc app=leetcode id=743 lang=golang
 *
 * [743] Network Delay Time
 */

// @lc code=start
func networkDelayTime(times [][]int, n int, k int) int {
	const inf int = 200
	g := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		g[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			g[i][j] = inf
		}
	}
	for _, t := range times {
		u, v, w := t[0], t[1], t[2]
		g[u][v] = w
	}

	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = inf
	}
	dist[k] = 0
	visited := make([]bool, n+1)
	for i := 0; i < n; i++ {
		x := -1
		for y := 1; y <= n; y++ {
			if !visited[y] && (x == -1 || dist[x] > dist[y]) {
				x = y
			}
		}
		visited[x] = true

		for nei, w := range g[x] {
			if dist[nei] > dist[x]+w {
				dist[nei] = dist[x] + w
			}
		}
	}

	max := 0
	for _, d := range dist {
		if d == inf {
			return -1
		}
		if max < d {
			max = d
		}
	}

	return max
}

// @lc code=end
