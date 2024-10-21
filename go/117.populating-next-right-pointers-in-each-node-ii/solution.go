package populatingnextrightpointersineachnodeii

/*
 * @lc app=leetcode id=117 lang=golang
 *
 * [117] Populating Next Right Pointers in Each Node II
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	type qItem struct {
		*Node
		lvl int
	}
	q := []qItem{{Node: root, lvl: 0}}
	for len(q) > 0 {
		qi := q[0]
		q = q[1:]

		if len(q) > 0 && q[0].lvl == qi.lvl {
			qi.Next = q[0].Node
		}
		if qi.Left != nil {
			q = append(q, qItem{Node: qi.Left, lvl: qi.lvl + 1})
		}
		if qi.Right != nil {
			q = append(q, qItem{Node: qi.Right, lvl: qi.lvl + 1})
		}
	}

	return root
}

// @lc code=end
