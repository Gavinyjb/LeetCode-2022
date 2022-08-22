package Problem_0101_SymmetricTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return help(root.Left, root.Right)
}
func help(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && help(left.Left, right.Right) && help(left.Right, right.Left)
}
