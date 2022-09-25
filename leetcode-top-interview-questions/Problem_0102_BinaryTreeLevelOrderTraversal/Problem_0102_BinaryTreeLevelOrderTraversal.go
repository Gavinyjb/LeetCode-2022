package Problem_0102_BinaryTreeLevelOrderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		ans := make([]int, 0)
		l := len(queue)

		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			ans = append(ans, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ret = append(ret, ans)
	}
	return ret
}

//func levelOrder(root *TreeNode) [][]int {
//	if root == nil {
//		return nil
//	}
//	ret := make([][]int, 0)
//	queue := make([]*TreeNode, 0)
//	queue = append(queue, root)
//	for len(queue) > 0 {
//		l := len(queue)
//		list := make([]int, 0)
//		for i := 0; i < l; i++ {
//			node := queue[0]
//			queue = queue[1:]
//			if node.Left != nil {
//				queue = append(queue, node.Left)
//			}
//			if node.Right != nil {
//				queue = append(queue, node.Right)
//			}
//			list = append(list, node.Val)
//		}
//		ret = append(ret, list)
//	}
//	return ret
//}

// Ints2TreeNode 利用 []int 生成 *TreeNode
func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != -1<<63 {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != -1<<63 {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
