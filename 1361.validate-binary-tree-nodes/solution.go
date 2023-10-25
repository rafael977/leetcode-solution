package validatebinarytreenodes

/*
 * @lc app=leetcode id=1361 lang=golang
 *
 * [1361] Validate Binary Tree Nodes
 */

// @lc code=start
func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	indeg := make([]int, n)
	for i := range leftChild {
		if l := leftChild[i]; l != -1 {
			indeg[l]++
		}
		if r := rightChild[i]; r != -1 {
			indeg[r]++
		}
	}
	root := 0
	for i, v := range indeg {
		if v == 0 {
			root = i
			break
		}
	}

	visited := make([]bool, n)
	q := []int{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if visited[node] {
			return false
		}
		visited[node] = true

		if l := leftChild[node]; l != -1 {
			q = append(q, l)
		}
		if r := rightChild[node]; r != -1 {
			q = append(q, r)
		}
	}

	allVisited := true
	for _, v := range visited {
		allVisited = allVisited && v
	}
	return allVisited
}

// @lc code=end
