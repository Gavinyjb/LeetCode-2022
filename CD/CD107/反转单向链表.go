package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func reverseList(head *Node) *Node {
	if head == nil {
		return nil
	}
	var pre *Node
	for head != nil {
		temp := head.Next
		head.Next = pre
		pre = head
		head = temp
	}
	return pre
}

type DoubleNode struct {
	Value int
	Last  *DoubleNode
	Next  *DoubleNode
}

func reverseDoubleList(head *DoubleNode) *DoubleNode {
	if head == nil {
		return head
	}
	var pre *DoubleNode
	for head != nil {
		temp := head.Next
		head.Next = pre
		head.Last = temp
		pre = head
		head = temp
	}
	return pre
}

func main() {
	n := 0
	fmt.Scan(&n)
	nodeList := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nodeList[i])
	}
	
	m := 0
	fmt.Scan(&m)
	doubleNodeList := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&doubleNodeList[i])
	}
	head1:=&Node{Value: nodeList[0],Next: nil}
	temp:=head1
	for i := 1; i < n; i++ {
		temp.Next=&Node{Value: nodeList[i]}
		temp=temp.Next
	}
	head2:=&DoubleNode{Value: doubleNodeList[0],Next: nil}
	temp2:=head2
	for i := 1; i < m; i++ {
		temp2.Next=&DoubleNode{Value: doubleNodeList[i]}
		last:=temp2
		temp2=temp2.Next
		temp2.Last=last
	}
	head1=reverseList(head1)
	head2=reverseDoubleList(head2)
	try:=head1
	for try!=nil{
		fmt.Print(try.Value)
		fmt.Print(" ")
		try=try.Next
	}
	fmt.Println()
	try2:=head2
	for try2!=nil{
		fmt.Print(try2.Value)
		fmt.Print(" ")
		try2=try2.Next
	}
}