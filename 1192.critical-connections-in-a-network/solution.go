package criticalconnectionsinanetwork

/*
 * @lc app=leetcode id=1192 lang=golang
 *
 * [1192] Critical Connections in a Network
 */

// @lc code=start
func criticalConnections(n int, connections [][]int) (ans [][]int) {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, 0)
	}
	for _, e := range connections {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}

	low := make([]int, n)
	dfn := make(map[int]int)
	i := 0

	var dfs func(u, p int)
	dfs = func(u, p int) {
		low[u] = i
		dfn[u] = i
		i++

		for _, v := range g[u] {
			if v == p {
				continue
			}

			vi, visited := dfn[v]
			if visited {
				low[u] = min(low[u], vi)
			} else {
				dfs(v, u)
				low[u] = min(low[u], low[v])
				if dfn[u] < low[v] {
					ans = append(ans, []int{u, v})
				}
			}
		}
	}

	dfs(0, -1)
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
