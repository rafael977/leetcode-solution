package main

type Node struct {
	Val       int
	Neighbors []*Node
}

/*
 * @lc app=leetcode id=133 lang=golang
 *
 * [133] Clone Graph
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	var nodes map[int]*Node = make(map[int]*Node)
	var cloneNode func(*Node) *Node
	cloneNode = func(node *Node) *Node {
		if nodes[node.Val] != nil {
			return nodes[node.Val]
		}

		ans := &Node{Val: node.Val, Neighbors: []*Node{}}
		nodes[node.Val] = ans
		for _, n := range node.Neighbors {
			cn := cloneNode(n)
			ans.Neighbors = append(ans.Neighbors, cn)
		}

		return ans
	}

	return cloneNode(node)
}

// @lc code=end
