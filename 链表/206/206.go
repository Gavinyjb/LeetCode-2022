package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	prev:=new(ListNode)
	for head!=nil{
		temp:=head.Next
		head.Next=prev
		prev=head
		head=temp
	}
	return prev
}

