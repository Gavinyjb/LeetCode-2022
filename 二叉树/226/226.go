package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归 自顶向下  先序
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree1(root.Left)
	invertTree1(root.Right)
	return root
}

//递归 自底向上 后序
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree2(root.Left)
	invertTree2(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

//迭代 层序
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		leave := queue[0]
		queue = queue[1:]
		leave.Left, leave.Right = leave.Right, leave.Left
		if leave.Left != nil {
			queue = append(queue, leave.Left)
		}
		if leave.Right != nil {
			queue = append(queue, leave.Right)
		}
	}
	return root
}
func main() {

}
