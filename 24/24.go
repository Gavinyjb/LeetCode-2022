package main

type ListNode struct {
	Val  int
	Next *ListNode
}
func swapPairs(head *ListNode) *ListNode {
	if head==nil {
		return head
	}
	//虚拟头结点
	dummyNode :=&ListNode{
		Next: head,
	}
	curr:=dummyNode
	for curr.Next!=nil&& curr.Next.Next!=nil{
		curr.Next=head.Next
		next:=head.Next.Next
		head.Next.Next=head
		head.Next=next
		curr=head
		head=next
	}
	return curr.Next

}
func main() {

}
