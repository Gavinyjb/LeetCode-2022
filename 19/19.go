package main

import "github.com/Gavinyjb/LeetCode-2022/01-structures"

//type structures.ListNode struct {
//	Val  int
//	Next *structures.ListNode
//}

func removeNthFromEnd(head *structures.ListNode, n int) *structures.ListNode {

	curr := &structures.ListNode{Next: head}
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
