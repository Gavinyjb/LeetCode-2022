package Problem_0092_ReserveBetween

func reserve(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
func reverseBetween2(head *ListNode, left, right int) *ListNode {
	if head == nil || left == right {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	oneEnd := pre
	twoStart := pre.Next
	twoEnd := oneEnd
	for i := 0; i < right-left+1; i++ {
		twoEnd = twoEnd.Next
	}

	threeStart := twoEnd.Next

	oneEnd.Next = nil
	twoEnd.Next = nil
	newTwoStart := reserve(twoStart)

	oneEnd.Next = newTwoStart
	twoStart.Next = threeStart
	return dummy.Next
}
