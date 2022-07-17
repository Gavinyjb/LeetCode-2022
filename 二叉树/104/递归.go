package main

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
