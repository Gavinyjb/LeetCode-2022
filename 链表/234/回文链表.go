package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	mid := findMiddle(head)
	tail := reverse(mid.Next)
	mid.Next = nil
	for head != nil && tail != nil {
		fmt.Printf("%v %v", head.Val, tail.Val)
		if head.Val != tail.Val {
			return false
		}
		head = head.Next
		tail = tail.Next
	}
	return true
}
func reverse(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	//不可初始化
	var prev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}
func findMiddle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
