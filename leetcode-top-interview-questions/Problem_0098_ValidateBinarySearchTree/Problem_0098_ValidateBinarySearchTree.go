package Problem_0098_ValidateBinarySearchTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return morris(root)
}
func morris(root *TreeNode) bool {
	if root == nil {
		return true
	}
	ret := true
	last := make([]int, 0)
	cur := root
	for cur != nil {
		if cur.Left == nil { //向右移动
			if len(last) > 0 {
				if last[0] >= cur.Val {
					ret = false
					last[0] = cur.Val
				} else {
					last[0] = cur.Val
				}
			} else {
				last = append(last, cur.Val)
			}
			cur = cur.Right
		} else {
			leftMostRight := cur.Left
			for leftMostRight.Right != nil && leftMostRight.Right != cur {
				leftMostRight = leftMostRight.Right
			}
			if leftMostRight.Right == nil {
				//第一次来
				leftMostRight.Right = cur
				cur = cur.Left
			} else {
				//第二次来
				if len(last) > 0 {
					if last[0] >= cur.Val {
						ret = false
						last[0] = cur.Val
					} else {
						last[0] = cur.Val
					}
				} else {
					last = append(last, cur.Val)
				}
				leftMostRight.Right = nil
				cur = cur.Right
			}
		}
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
