package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
func isSymmetric1(root *TreeNode) bool {
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

//迭代

func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root.Left, root.Right)

	for len(queue) != 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue //很重要！！！！！！！！
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Right) // 加入左节点左孩子  加入右节点右孩子
		queue = append(queue, left.Right, right.Left) // 加入左节点右孩子  加入左节点右孩子
	}
	return true
}

//递归 加翻转
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := invertTree2(root.Left)
	return isSameTree(left, root.Right)
}
func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(q.Left, p.Left) && isSameTree(q.Right, p.Right)
}

//  翻转二叉树
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree2(root.Left)
	invertTree2(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
func main() {

}
