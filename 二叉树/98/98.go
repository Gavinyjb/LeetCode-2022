package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	res := inorder(root)
	if len(res) <= 1 {
		return true
	}
	for i := 1; i < len(res); i++ {
		if res[i] < res[i-1] {
			return false
		}
	}
	return true
}
func inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := make([]int, 0)
	left := inorder(root.Left)
	ret = append(ret, left...)
	ret = append(ret, root.Val)
	right := inorder(root.Right)
	ret = append(ret, right...)
	return ret
}
