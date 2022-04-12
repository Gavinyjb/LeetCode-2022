package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	//虚拟头结点
	dummyNode := &ListNode{
		Next: head,
	}
	curr := dummyNode
	for curr.Next != nil && curr.Next.Next != nil {
		temp1 := curr.Next
		temp3 := curr.Next.Next.Next
		curr.Next = curr.Next.Next
		curr.Next.Next = temp1
		curr.Next.Next.Next = temp3
		curr = curr.Next.Next
	}
	return dummyNode.Next
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	for curr := dummy; curr != nil && curr.Next != nil && curr.Next.Next != nil; {
		curr, curr.Next, curr.Next.Next, curr.Next.Next.Next = curr.Next, curr.Next.Next, curr.Next.Next.Next, curr.Next
	}
	return dummy.Next
}
func main() {
	a, b, c, d := 1, 2, 3, 4
	a, b, c, d = b, c, d, b
	fmt.Printf("a:%v,b:%v,c:%v,d:%d", a, b, c, d)

}
