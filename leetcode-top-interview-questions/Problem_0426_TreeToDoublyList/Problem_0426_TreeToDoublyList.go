package Problem_0426_TreeToDoublyList

import (
	"fmt"
)

//Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return root
	}
	var head, end *Node
	stack := make([]*Node, 0)
	var lastNode *Node
	var minNode *Node

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if minNode == nil {
			minNode = node
		} else {
			if node.Val < minNode.Val {
				minNode = node
			}
		}

		if lastNode == nil {
			head = node
		} else {
			lastNode.Right = node
			node.Left = lastNode
		}
		lastNode = node
		root = node.Right
	}
	if lastNode != nil && head != nil {
		lastNode.Right = head
		head.Left = lastNode
	}
	return minNode
}
