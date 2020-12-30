package main

import (
	"fmt"
	"math"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func maxPathSum(root *TreeNode) int {
	curMax, _ := helper(root)

	return curMax
}

func helper(root *TreeNode) (curMax, pathMax int) {
	if root == nil {
		return math.MinInt64, 0
	}

	lcm, lpm := helper(root.Left)
	rcm, rpm := helper(root.Right)

	curMax = max([]int{lcm, rcm, root.Val, lpm + root.Val, rpm + root.Val, lpm + rpm + root.Val})
	pathMax = max([]int{root.Val, lpm + root.Val, rpm + root.Val})

	return curMax, pathMax
}

func max(nums []int) int {
	r := nums[0]
	for _, v := range nums[1:] {
		if v > r {
			r = v
		}
	}

	return r
}

func main() {
	root := BuildTree("1,2,3")
	fmt.Println(maxPathSum(root))

	root = BuildTree("-10,9,20,null,null,15,7")
	fmt.Println(maxPathSum(root))

	root = BuildTree("1,2,3,-1,-1,4,-1")
	fmt.Println(maxPathSum(root))

	root = BuildTree("-3")
	fmt.Println(maxPathSum(root))
}
