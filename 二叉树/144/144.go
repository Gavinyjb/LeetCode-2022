package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	output := make([]int, 0)
	proorder(root, &output)
	return output
}
func proorder(root *TreeNode, output *[]int) {
	if root != nil {
		*output = append(*output, root.Val)
		proorder(root.Left, output)
		proorder(root.Right, output)
	}
}
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := make([]int, 0)
	ret = append(ret, root.Val)
	left := preorderTraversal(root.Left)
	right := preorderTraversal(root.Right)
	ret = append(ret, left...)
	ret = append(ret, right...)
	return ret

}

//迭代
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			// 前序遍历，所以先保存结果
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left //一直向左
		}
		//pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}
func main() {

}
