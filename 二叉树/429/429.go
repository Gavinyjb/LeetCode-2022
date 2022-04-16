package main

//Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		l := len(queue)
		list := make([]int, 0)
		for i := 0; i < l; i++ {
			leave := queue[0]
			queue = queue[1:]
			list = append(list, leave.Val)
			for j := 0; j < len(leave.Children); j++ {
				queue = append(queue, leave.Children[j])
			}
		}
		res = append(res, list)
	}
	return res
}
func main() {

}
