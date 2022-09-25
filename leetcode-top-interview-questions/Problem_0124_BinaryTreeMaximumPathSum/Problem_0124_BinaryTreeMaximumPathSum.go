package Problem_0124_BinaryTreeMaximumPathSum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Info struct {
	maxPathSum      int
	maxPathFromHead int
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return help(root).maxPathSum
}
func help(node *TreeNode) Info {
	if node == nil {
		return Info{
			maxPathSum:      math.MinInt32,
			maxPathFromHead: math.MinInt32,
		}
	}
	leftInfo := help(node.Left)
	rightInfo := help(node.Right)
	//左边
	p1 := leftInfo.maxPathFromHead + node.Val
	//右边
	p2 := rightInfo.maxPathFromHead + node.Val
	//
	p3 := max(leftInfo.maxPathSum, rightInfo.maxPathSum)
	//加上左右两边
	p4 := node.Val + leftInfo.maxPathFromHead + rightInfo.maxPathFromHead
	maxPathFromHead := max(node.Val, p1)
	maxPathFromHead = max(maxPathFromHead, p2)

	maxPathSum := max(maxPathFromHead, p3)
	maxPathSum = max(maxPathSum, p4)

	return Info{
		maxPathSum:      maxPathSum,
		maxPathFromHead: maxPathFromHead,
	}
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
