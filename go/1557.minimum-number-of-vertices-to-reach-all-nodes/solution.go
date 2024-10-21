package minimumnumberofverticestoreachallnodes

/*
 * @lc app=leetcode id=1557 lang=golang
 *
 * [1557] Minimum Number of Vertices to Reach All Nodes
 */

// @lc code=start
func findSmallestSetOfVertices(n int, edges [][]int) (ans []int) {
	hasIn := make(map[int]bool)
	for _, e := range edges {
		hasIn[e[1]] = true
	}

	for i := 0; i < n; i++ {
		if !hasIn[i] {
			ans = append(ans, i)
		}
	}
	return
}

// @lc code=end
