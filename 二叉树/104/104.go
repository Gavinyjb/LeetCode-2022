package main

import gavin "github.com/Gavinyjb/LeetCode-2022/01-structures"

type TreeNode = gavin.TreeNode

//递归 后序  其实是求根节点高度的逻辑
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max(left, right) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//递归 先序 求节点深度
func maxDepth(root *TreeNode) int {
	result := 0
	if root == nil {
		return result
	}
	var getDepth func(node *TreeNode, depth int)
	getDepth = func(node *TreeNode, depth int) {
		if depth > result {
			result = depth
		}
		if node.Left == nil && node.Right == nil { //到达叶子节点
			return
		}
		if node.Left != nil {
			depth++
			getDepth(node.Left, depth)
			depth--
		}
		if node.Right != nil {
			depth++
			getDepth(node.Right, depth)
			depth--
		}
	}
	getDepth(root, 1)
	return result
}

func main() {

}
