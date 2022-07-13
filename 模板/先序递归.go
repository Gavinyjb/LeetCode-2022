package 模板

import "fmt"

func preorder(root *TreeNode) {
	if root == nil {
		return
	}
	//先根后左右
	fmt.Println(root.Val)
	preorder(root.Left)
	preorder(root.Right)
}
