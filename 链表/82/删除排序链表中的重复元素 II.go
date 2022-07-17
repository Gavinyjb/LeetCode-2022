package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	//头结点可能被删除 需要添加虚拟节点
	if head == nil {
		return nil
	}
	dummy := &ListNode{
		Val: 0,
	}
	dummy.Next = head
	current := dummy
	var rmval int
	for current.Next != nil && current.Next.Next != nil {
		if current.Next.Val == current.Next.Next.Val {
			rmval = current.Next.Val
			for current.Next.Next != nil && rmval == current.Next.Next.Val {
				current.Next = current.Next.Next
			}
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return dummy.Next
}
