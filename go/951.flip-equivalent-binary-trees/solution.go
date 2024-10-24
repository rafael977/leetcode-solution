// Created by Rafael Shen at 2024/10/24 10:51
// leetgo: 1.4.9
// https://leetcode.com/problems/flip-equivalent-binary-trees/

package flipequivalentbinarytrees

import . "github.com/rafael977/leetcode-solution/datastruct"

// @lc code=begin

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	} else if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}

	return (flipEquiv(root1.Left, root2.Left) &&
		flipEquiv(root1.Right, root2.Right)) ||
		(flipEquiv(root1.Left, root2.Right) &&
			flipEquiv(root1.Right, root2.Left))
}

// @lc code=end
