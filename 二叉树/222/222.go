package main

import (
	gavin "github.com/Gavinyjb/LeetCode-2022/01-structures"
)

type TreeNode = gavin.TreeNode

//递归
func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

//迭代
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		l := q[0]
		q = q[1:]
		count++
		if l.Left != nil {
			q = append(q, l.Left)
		}
		if l.Right != nil {
			q = append(q, l.Right)
		}
	}
	return count
}
func main() {

}
