// Created by Rafael Shen at 2024/10/22 21:03
// leetgo: 1.4.9
// https://leetcode.com/problems/kth-largest-sum-in-a-binary-tree/

package main

import (
	"sort"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

// @lc code=begin

func kthLargestLevelSum(root *TreeNode, k int) (ans int64) {
	sums := make([]int, 0)
	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) > 0 {
		nq := make([]*TreeNode, 0)
		s := 0
		for _, v := range q {
			s += v.Val
			if v.Left != nil {
				nq = append(nq, v.Left)
			}
			if v.Right != nil {
				nq = append(nq, v.Right)
			}
		}
		sums = append(sums, s)
		q = nq
	}

	if len(sums) < k {
		return -1
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))
	return int64(sums[k-1])
}

// @lc code=end
