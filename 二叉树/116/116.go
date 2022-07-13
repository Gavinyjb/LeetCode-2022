package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	levelOrder(root)
	return root
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
			if i < l-1 {
				leave.Next = queue[0]
			}
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

func main() {

}
