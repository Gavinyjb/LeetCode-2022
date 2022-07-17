package leetcode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var prev *ListNode
	for head != nil {
		//保存当前的head.Next状态
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}
