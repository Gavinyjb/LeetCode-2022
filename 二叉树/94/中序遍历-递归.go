package leetcdoe

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	left := inorderTraversal(root.Left)
	var result []int
	result = append(result, left...)
	result = append(result, root.Val)
	right := inorderTraversal(root.Right)
	result = append(result, right...)
	return result
}
