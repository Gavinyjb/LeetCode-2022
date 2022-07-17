package leetcode

//指针比较时直接比较对象，不要用值比较，链表中有可能存在重复值情况
//第一次相交后，快指针需要从下一个节点开始和头指针一起匀速移动
func detectCycle2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		if fast == slow {
			// 慢指针重新从头开始移动，快指针从第一次相交点下一个节点开始移动
			fast = fast.Next //注意
			slow = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return nil
}
