package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

//快慢指针 判断 fast 及 fast.Next 是否为 nil 值
//递归 mergeSort 需要断开中间节点
//递归返回条件为 head 为 nil 或者 head.Next 为 nil

func sortList(head *ListNode) *ListNode {
	//归并排序，找中点和合并操作
	return mergeSort(head)
}
func mergeSort(head *ListNode) *ListNode {
	// 如果只有一个节点 直接就返回这个节点
	if head == nil || head.Next == nil {
		return head
	}
	mid := FindMiddle(head)
	// fast如果初始化为head.Next则中点在slow.Next
	// fast初始化为head,则中点在slow
	//断开中间节点
	tail := mid.Next
	mid.Next = nil
	left := mergeSort(head)
	right := mergeSort(tail)
	result := mergeTwoList(left, right)
	return result
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}
	if l1 != nil {
		head.Next = l1
	} else {
		head.Next = l2
	}
	return dummy.Next
}
func FindMiddle(head *ListNode) *ListNode {
	slow := head
	fast := head.Next
	//快指针先为nil
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
