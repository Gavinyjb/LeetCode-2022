package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, isBalance := maxDepth(root)
	return isBalance
}
func maxDepth(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	left, boolLeft := maxDepth(root.Left)
	right, boolRight := maxDepth(root.Right)
	if boolLeft == false || boolRight == false || left-right > 1 || right-left > 1 {
		return -1, false
	}
	if left > right {
		return left + 1, true
	} else {
		return right + 1, true
	}
}

//func maxDepth(root *TreeNode) int {
//	// check
//	if root == nil {
//		return 0
//	}
//	left := maxDepth(root.Left)
//	right := maxDepth(root.Right)
//	// 为什么返回-1呢？（变量具有二义性）
//	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
//		return -1
//	}
//	if left > right {
//		return left + 1
//	}
//	return right + 1
//}

func main() {

}
