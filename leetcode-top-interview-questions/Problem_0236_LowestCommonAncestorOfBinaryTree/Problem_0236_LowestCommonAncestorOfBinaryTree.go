package Problem_0236_LowestCommonAncestorOfBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return process(root, p, q).ans
}

type info struct {
	findP bool
	findQ bool
	ans   *TreeNode
}

func process(root, p, q *TreeNode) info {
	if root == nil {
		return info{
			findP: false,
			findQ: false,
			ans:   nil,
		}
	}
	leftInfo := process(root.Left, p, q)
	rightInfo := process(root.Right, p, q)
	findP := root == p || leftInfo.findP || rightInfo.findP
	findQ := root == q || leftInfo.findQ || rightInfo.findQ
	var ans *TreeNode
	if leftInfo.ans != nil {
		ans = leftInfo.ans
	} else if rightInfo.ans != nil {
		ans = rightInfo.ans
	} else {
		if findP && findQ {
			ans = root
		}
	}
	return info{findP: findP, findQ: findQ, ans: ans}
}
