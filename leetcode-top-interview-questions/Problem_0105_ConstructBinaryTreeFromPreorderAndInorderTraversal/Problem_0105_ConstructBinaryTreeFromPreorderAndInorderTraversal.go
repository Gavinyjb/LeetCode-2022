package Problem_0105_ConstructBinaryTreeFromPreorderAndInorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	root := new(TreeNode)
	root.Val = preorder[0]
	for rootIdx, node := range inorder {
		if node == root.Val {
			root.Left = buildTree(preorder[1:rootIdx+1], inorder[:rootIdx])
			root.Right = buildTree(preorder[rootIdx+1:], inorder[rootIdx+1:])
		}
	}
	return root
}
