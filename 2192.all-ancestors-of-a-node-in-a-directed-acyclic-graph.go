package leetcodesolution

/*
 * @lc app=leetcode id=2192 lang=golang
 *
 * [2192] All Ancestors of a Node in a Directed Acyclic Graph
 */

// @lc code=start
func getAncestors(n int, edges [][]int) (ans [][]int) {
	g := make([][]int, n)
	for _, e := range edges {
		g[e[1]] = append(g[e[1]], e[0])
	}

	vis := make([]bool, n)
	var dfs func(x int)
	dfs = func(x int) {
		vis[x] = true
		for _, v := range g[x] {
			if !vis[v] {
				dfs(v)
			}
		}
	}
	ans = make([][]int, n)
	for i := range ans {
		clear(vis)
		dfs(i)
		vis[i] = false
		for x, v := range vis {
			if v {
				ans[i] = append(ans[i], x)
			}
		}
	}

	return
}

// @lc code=end
