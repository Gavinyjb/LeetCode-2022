package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//层序遍历
func minDepth1(root *TreeNode) int {
	return levelOrder(root)
}

func levelOrder(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	ret := 0
	for len(queue) != 0 {
		l := len(queue)
		ret++
		for i := 0; i < l; i++ {
			leave := queue[0]
			queue = queue[1:]
			if leave.Left == nil && leave.Right == nil {
				//叶子节点
				return ret
			}
			if leave.Left != nil {
				queue = append(queue, leave.Left)
			}
			if leave.Right != nil {
				queue = append(queue, leave.Right)
			}
		}
	}
	return ret
}

//DFS 深度搜索 递归
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minDep := math.MaxInt64
	if root.Left != nil {
		minDep = min(minDepth(root.Left), minDep)
	}
	if root.Right != nil {
		minDep = min(minDepth(root.Right), minDep)
	}
	return minDep + 1
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {

}
