package Problem_0144_BinaryTreePreOrderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func morrisPre(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := make([]int, 0)
	cur := root
	for cur != nil {
		if cur.Left == nil {
			ret = append(ret, cur.Val)
			cur = cur.Right
		} else {
			leftMostRight := cur.Left
			for leftMostRight.Right != nil && leftMostRight.Right != cur {
				leftMostRight = leftMostRight.Right
			}
			if leftMostRight.Right == nil { //第一次来
				ret = append(ret, cur.Val)
				leftMostRight.Right = cur
				cur = cur.Left
			} else { //第二次来
				leftMostRight.Right = nil
				cur = cur.Right
			}
		}
	}
	return ret
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	ret := make([]int, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			ret = append(ret, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return ret
}

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
