package main

import "math"

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
func largestValues(root *TreeNode) []int {
	level := levelOrder(root)
	ret := make([]int, 0)
	for _, ints := range level {
		max := math.MinInt64
		for _, v := range ints {
			if v > max {
				max = v
			}
		}
		ret = append(ret, max)
	}
	return ret
}

func main() {

}
