package main

import "github.com/Gavinyjb/LeetCode-2022/01-structures"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	curr := &ListNode{Next: head}
	slow, fast := curr, curr
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return curr.Next
}
func main() {

}
