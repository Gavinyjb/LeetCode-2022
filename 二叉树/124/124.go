package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt
	oneSideMax(root, &res)
	return res
}

func oneSideMax(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	leftmaxSum := max(0, oneSideMax(root.Left, res))
	rightmaxSum := max(0, oneSideMax(root.Right, res))
	pathMax := root.Val + leftmaxSum + rightmaxSum
	*res = max(*res, pathMax)
	return max(leftmaxSum, rightmaxSum) + root.Val
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
