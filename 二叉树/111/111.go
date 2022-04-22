package main

import (
	gavin "github.com/Gavinyjb/LeetCode-2022/01-structures"
)

type TreeNode = gavin.TreeNode

//递归 后序
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return 1 + minDepth(root.Right)
	}
	if root.Right == nil {
		return 1 + minDepth(root.Left)
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	return 1 + min(left, right)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//层序
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		l := len(queue)
		depth++
		for i := 0; i < l; i++ {
			leave := queue[0]
			queue = queue[1:]
			if leave.Left == nil && leave.Right == nil { //叶子节点
				return depth
			}
			if leave.Left != nil {
				queue = append(queue, leave.Left)
			}
			if leave.Right != nil {
				queue = append(queue, leave.Right)
			}
		}
	}
	return depth
}

func main() {

}
