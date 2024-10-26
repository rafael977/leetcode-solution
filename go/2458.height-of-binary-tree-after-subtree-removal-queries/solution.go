// Created by Rafael Shen at 2024/10/26 09:37
// leetgo: 1.4.9
// https://leetcode.com/problems/height-of-binary-tree-after-subtree-removal-queries/

package heightofbinarytreeaftersubtreeremovalqueries

import . "github.com/rafael977/leetcode-solution/datastruct"

// @lc code=begin

func treeQueries(root *TreeNode, queries []int) (ans []int) {
	height := map[*TreeNode]int{}

	var getHeight func(*TreeNode) int

	getHeight = func(node *TreeNode) int {

		if node == nil {

			return 0

		}

		height[node] = 1 + max(getHeight(node.Left), getHeight(node.Right))

		return height[node]

	}

	getHeight(root)

	res := make([]int, len(height)+1)

	var dfs func(*TreeNode, int, int)

	dfs = func(node *TreeNode, depth, restH int) {

		if node == nil {

			return

		}

		depth++

		res[node.Val] = restH

		dfs(node.Left, depth, max(restH, depth+height[node.Right]))

		dfs(node.Right, depth, max(restH, depth+height[node.Left]))

	}

	dfs(root, -1, 0)

	for i, q := range queries {

		queries[i] = res[q]

	}

	return queries
}

// @lc code=end
