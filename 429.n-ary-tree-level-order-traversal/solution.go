package narytreelevelordertraversal

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=429 lang=golang
 *
 * [429] N-ary Tree Level Order Traversal
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) (ans [][]int) {
	if root == nil {
		return
	}

	q := []*Node{root}
	for len(q) > 0 {
		vals := make([]int, 0)
		newQ := make([]*Node, 0)
		for _, n := range q {
			vals = append(vals, n.Val)
			newQ = append(newQ, n.Children...)
		}

		ans = append(ans, vals)
		q = newQ
	}
	return
}

// @lc code=end

type Node struct {
	Val      int
	Children []*Node
}

func buildNTree(in string) *Node {
	in = strings.TrimSpace(in)
	if len(in) == 0 {
		return nil
	}

	parts := strings.Split(in, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	item := parts[0]
	itemVal, _ := strconv.Atoi(item)
	root := &Node{
		Val: itemVal,
	}

	q := []*Node{root}
	i := 1
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		i++
		for i < len(parts) && parts[i] != "null" {
			itemVal, _ := strconv.Atoi(parts[i])
			itemNode := &Node{Val: itemVal}
			node.Children = append(node.Children, itemNode)
			q = append(q, itemNode)
			i++
		}
	}

	return root
}
