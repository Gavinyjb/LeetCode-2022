package 模板

func MorrisPre(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var mostRight *TreeNode
	cur := root
	ret := make([]int, 0)
	for cur != nil {
		//如果当前节点没有左子树
		if cur.Left == nil { //当前节点向右移动(只有没有左子树的节点会走到这一步)
			ret = append(ret, cur.Val)
			cur = cur.Right
		} else { //找到当前节点的左子树的最右节点
			mostRight = cur.Left

			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			//此时mostRight就是当前节点左子树的最右节点
			if mostRight.Right == nil { //说明是第一次来到
				ret = append(ret, cur.Val)
				mostRight.Right = cur //让其指向当前节点
				cur = cur.Left        //cur向左移动
			} else { //说明是第二次来到(只有有左子树的时候才会有第二次来到)
				mostRight.Right = nil
				cur = cur.Right
			}
		}
	}
	return ret
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

//Morris后序遍历，自己写的参考左神
//打印时机就放到可以回到自己并且第二次回到的时候，但是不是打印自己而是打印左树的右边界，并且是逆序打印
//但是这样的方式会错过整棵树的右边界，所以最后自己单独打印一下右边界
//因为一棵树是这么拆分下来的：
func MyMorrisPostOrderTraverse(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var mostRight *TreeNode
	cur := root
	var ret []int
	for cur != nil {
		if cur.Left == nil {
			cur = cur.Right
		} else {
			mostRight = cur.Left
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil { //说明是第一次遍历你到
				mostRight.Right, cur = cur, cur.Left
			} else {
				mostRight.Right = nil
				//打印左子树的右边界
				ret = append(ret, prinSubTreeRightEdge(cur.Left)...)
				cur = cur.Right
			}
		}
	}

	//最后再打印一下整棵树的右边界即可
	ret = append(ret, prinSubTreeRightEdge(root.Right)...)
	return ret
}

func prinSubTreeRightEdge(head *TreeNode) []int {
	if head == nil {
		return nil
	}

	var ret []int

	//反转
	var prev *TreeNode
	cur := head
	for cur != nil {
		prev, cur, cur.Right = cur, cur.Right, prev
	}

	//此时prev指向反转之后的头
	//开始打印
	cur = prev
	for cur != nil {
		ret = append(ret, cur.Val)
		cur = cur.Right
	}

	//反转回去
	prev, cur = nil, prev
	for cur != nil {
		prev, cur, cur.Right = cur, cur.Right, prev
	}

	return ret
}
