package binarysearchtreeiterator

import (
	. "github.com/rafael977/leetcode-solution/datastruct"
)

/*
 * @lc app=leetcode id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	nodes []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{
		nodes: make([]*TreeNode, 0),
	}
	for root != nil {
		it.nodes = append(it.nodes, root)
		root = root.Left
	}
	return it
}

func (this *BSTIterator) Next() int {
	n := len(this.nodes)
	node := this.nodes[n-1]
	this.nodes = this.nodes[:n-1]
	it := node.Right
	for it != nil {
		this.nodes = append(this.nodes, it)
		it = it.Left
	}
	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.nodes) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end
