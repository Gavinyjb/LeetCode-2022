package main

type ListNode struct {
	Val  int
	Next *ListNode
}
//双指针
func reverseList1(head *ListNode) *ListNode {
	if head==nil{
		return head
	}
	//创建虚拟头部
	var behind *ListNode
	curr:=head
	for curr!=nil{
		next:=curr.Next
		curr.Next=behind
		behind=curr
		curr=next
	}
	return behind
}

func reverseList(head *ListNode) *ListNode {
	if head==nil{
		return head
	}
	return reverse(nil,head)
}
func reverse(pre,head *ListNode)*ListNode{
	if head==nil{
		return pre
	}
	next:=head.Next
	head.Next=pre
	return  reverse(head,next)

}
func main() {

}
