package lowestcommonancestorofabinarysearchtree

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=235 lang=golang
 *
 * [235] Lowest Common Ancestor of a Binary Search Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}

	for root != nil {
		if root.Val >= p.Val && root.Val <= q.Val {
			break
		}
		if root.Val < p.Val {
			root = root.Right
		} else if root.Val > q.Val {
			root = root.Left
		}
	}

	return root
}

// @lc code=end
