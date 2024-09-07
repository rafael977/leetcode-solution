package linkedlistinbinarytree

import (
	. "github.com/rafael977/leetcode-solution/datastruct"
)

/*
 * @lc app=leetcode id=1367 lang=golang
 *
 * [1367] Linked List in Binary Tree
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubPath(head *ListNode, root *TreeNode) bool {
	var dfs func(e *ListNode, n *TreeNode) bool
	dfs = func(e *ListNode, n *TreeNode) bool {
		if e == nil {
			return true
		}
		if n == nil || e.Val != n.Val {
			return false
		}
		return dfs(e.Next, n.Left) ||
			dfs(e.Next, n.Right)
	}

	if root == nil {
		return false
	}
	return dfs(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

// @lc code=end
