package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	return reverse(nil, head)
}
func reverse(prev, head *ListNode) *ListNode {
	if head == nil {
		return prev
	}
	//保存当前的head.Next状态
	temp := head.Next
	head.Next = prev
	return reverse(head, temp)
}

//func reverse(pre,head *ListNode)*ListNode{
//	if head==nil{
//		return pre
//	}
//	next:=head.Next
//	head.Next=pre
//	return  reverse(head,next)
//}
