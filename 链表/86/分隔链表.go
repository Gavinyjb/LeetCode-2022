package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	//思路：将大于x的节点，放到另外一个链表，最后连接这两个链表
	if head == nil {
		return head
	}
	small := &ListNode{}
	smallHead := small
	large := &ListNode{}
	largeHear := large
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	large.Next = nil
	small.Next = largeHear.Next
	return smallHead.Next
}
