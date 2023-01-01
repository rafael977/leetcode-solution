package allpathsfromsourcetotarget

/*
 * @lc app=leetcode id=797 lang=golang
 *
 * [797] All Paths From Source to Target
 */

// @lc code=start
func allPathsSourceTarget(graph [][]int) (ans [][]int) {
	n := len(graph)
	path := make([]int, 0)

	var dfs func(i int)
	dfs = func(i int) {
		path = append(path, i)
		defer func() {
			path = path[:len(path)-1]
		}()
		if i == n-1 {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for _, v := range graph[i] {
			dfs(v)
		}
	}
	dfs(0)
	return
}

// @lc code=end
