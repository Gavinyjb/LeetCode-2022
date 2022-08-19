package problem0019removenthnodefromendoflist
type ListNode struct{
	Val int
	Next *ListNode
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head==nil{
		return head
	}
	dummy:=new(ListNode)
	dummy.Next=head
	fast:=dummy
	slow:=dummy
	for i := n; i >0; i-- {
		fast=fast.Next
	}
	for fast.Next!=nil{
		fast=fast.Next
		slow=slow.Next
	}
	slow.Next=slow.Next.Next
	return dummy.Next
}