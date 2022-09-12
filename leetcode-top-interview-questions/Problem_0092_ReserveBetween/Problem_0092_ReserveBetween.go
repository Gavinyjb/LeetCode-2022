package Problem_0092_ReserveBetween

import "fmt"

//输入：head = [1,2,3,4,5], left = 2, right = 4
//输出：[1,4,3,2,5]
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dunmmy := &ListNode{}
	dunmmy.Next = head
	cur := head
	i := 1
	for ; i < left; i++ {
		cur = cur.Next
	}
	end1 := cur
	start := cur.Next
	for ; i < right; i++ {
		cur = cur.Next
	}
	end := cur.Next
	var start2 *ListNode
	if end != nil {
		start2 = end.Next
		end.Next = nil
	} else {
		start2 = nil
	}

	temp := start
	for temp != nil {
		fmt.Println(temp.Val)
		temp = temp.Next
	}
	newStart := reverse(start)
	cur = newStart
	end1.Next = newStart
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = start2
	return dunmmy.Next
}
func reverse(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pre := &ListNode{}

	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre.Next
		pre.Next = cur
		cur = temp
	}
	return pre
}
