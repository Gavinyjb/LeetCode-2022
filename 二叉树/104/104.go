package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//层序遍历
func maxDepth1(root *TreeNode) int {
	level := levelOrder(root)
	return len(level)
}
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		l := len(queue)
		list := make([]int, 0)
		for i := 0; i < l; i++ {
			leave := queue[0]
			queue = queue[1:]
			list = append(list, leave.Val)
			if leave.Left != nil {
				queue = append(queue, leave.Left)
			}
			if leave.Right != nil {
				queue = append(queue, leave.Right)
			}
		}
		res = append(res, list)
	}
	return res
}

//DFS 深度搜索 递归
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// depth 记录遍历到的节点的深度    ret:记录最大深度
	depth, ret := 0, 0
	var help func(root *TreeNode)
	help = func(root *TreeNode) {
		//到达叶子节点，更新最大深度
		if root == nil {
			if depth > ret {
				ret = depth
			}
		}
		depth++ //进入节点
		help(root.Left)
		help(root.Right)
		depth-- //离开节点
	}
	help(root)
	return ret
}

func main() {

}
