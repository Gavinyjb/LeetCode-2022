package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
func averageOfLevels(root *TreeNode) []float64 {
	level := levelOrder(root)
	ret := make([]float64, 0)
	for _, ints := range level {
		sum := 0
		for _, v := range ints {
			sum += v
		}
		ret = append(ret, float64(sum)/float64(len(ints)))
	}
	return ret
}

func main() {

}
