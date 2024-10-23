// Created by Rafael Shen at 2024/10/23 09:52
// leetgo: 1.4.9
// https://leetcode.com/problems/cousins-in-binary-tree-ii/

package cousinsinbinarytreeii

import . "github.com/rafael977/leetcode-solution/datastruct"

// @lc code=begin

func replaceValueInTree(root *TreeNode) (ans *TreeNode) {
	q := []*TreeNode{root}
	root.Val = 0
	for len(q) > 0 {
		nq := make([]*TreeNode, 0)
		sum := 0
		for _, v := range q {
			if v.Left != nil {
				sum += v.Left.Val
			}
			if v.Right != nil {
				sum += v.Right.Val
			}
		}
		for _, v := range q {
			csum := 0
			if v.Left != nil {
				csum += v.Left.Val
			}
			if v.Right != nil {
				csum += v.Right.Val
			}
			if v.Left != nil {
				v.Left.Val = sum - csum
				nq = append(nq, v.Left)
			}
			if v.Right != nil {
				v.Right.Val = sum - csum
				nq = append(nq, v.Right)
			}
		}

		q = nq
	}
	return root
}

// @lc code=end
