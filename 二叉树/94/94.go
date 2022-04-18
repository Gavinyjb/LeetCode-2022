package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := make([]int, 0)
	left := inorderTraversal(root.Left)
	ret = append(ret, left...)
	ret = append(ret, root.Val)
	right := inorderTraversal(root.Right)
	ret = append(ret, right...)
	return ret
}

// 思路：通过stack 保存已经访问的元素，用于原路返回

func inorderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		root = node.Right
	}
	return result
}
func main() {

}
