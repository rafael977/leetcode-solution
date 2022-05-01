package isgraphbipartite

/*
 * @lc app=leetcode id=785 lang=golang
 *
 * [785] Is Graph Bipartite?
 */

// @lc code=start
func isBipartite(graph [][]int) bool {
	color := make(map[int]int)
	valid := true
	dfs := func(i int) {
		color[i] = 1
		q := []int{i}
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			for _, v := range graph[u] {
				if color[v] == 0 {
					color[v] = -color[u]
					q = append(q, v)
				} else if color[v] == color[u] {
					valid = false
					return
				}
			}
		}
	}

	for i := 0; i < len(graph); i++ {
		if color[i] == 0 {
			dfs(i)
		}
	}

	return valid
}

// @lc code=end
