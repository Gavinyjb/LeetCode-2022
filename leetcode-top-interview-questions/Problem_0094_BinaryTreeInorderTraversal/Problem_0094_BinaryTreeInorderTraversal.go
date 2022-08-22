package Problem_0094_BinaryTreeInorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MorrisIn(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var mostRight *TreeNode
	cur := root
	ret := make([]int, 0)
	for cur != nil {
		//如果当前节点没有左子树
		if cur.Left == nil { //节点右移
			ret = append(ret, cur.Val)
			cur = cur.Right
		} else { //找到当晚节点的最右节点
			mostRight = cur.Left

			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			//此时的mostRight就是当前节点左子树的最右节点
			if mostRight.Right == nil { //第一次到
				mostRight.Right = cur
				cur = cur.Left
			} else { //第二次到达
				ret = append(ret, cur.Val)
				mostRight.Right = nil
				cur = cur.Right
			}
		}
	}
	return ret
}
func morrisIn(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := make([]int, 0)
	leftMostRight := new(TreeNode)
	cur := root
	for cur != nil {
		if cur.Left == nil { //左子树为空 右边移动
			ret = append(ret, cur.Val)
			cur = cur.Right
		} else {
			leftMostRight = cur.Left
			for leftMostRight.Right != nil && leftMostRight.Right != cur {
				leftMostRight = leftMostRight.Right
			}
			//此时的mostRight就是当前节点左子树的最右节点
			if leftMostRight.Right == nil { //第一次到达
				leftMostRight.Right = cur
				cur = cur.Left
			} else { //第二次到达
				ret = append(ret, cur.Val)
				leftMostRight.Right = nil
				cur = cur.Right
			}
		}
	}
	return ret
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.Val)
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
