package 模板

//func levelOrder(root *TreeNode) [][]int {
//	if root == nil {
//		return nil
//	}
//	//通过上一层的长度确定下一层的长度
//	res := make([][]int, 0)
//	queue := make([]*TreeNode, 0)
//	queue = append(queue, root)
//	for len(queue) > 0 {
//		list := make([]int, 0)
//		// 为什么要取length？
//		// 记录当前层有多少元素（遍历当前层，再添加下一层）
//		l := len(queue)
//		for i := 0; i < l; i++ {
//			//出队列
//			leave := queue[0]
//			queue = queue[1:]
//			list = append(list, leave.Val)
//			if leave.Left != nil {
//				queue = append(queue, leave.Left)
//			}
//			if leave.Right != nil {
//				queue = append(queue, leave.Right)
//			}
//		}
//		res = append(res, list)
//	}
//	return res
//}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for root != nil || len(queue) > 0 {
		list := make([]int, 0)
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, list)
	}
	return result
}
