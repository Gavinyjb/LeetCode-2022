/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 package leetcodetopinterviewquestions
 
 type ListNode struct{
	Val int
	Next *ListNode
 }
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1==nil{
		return l2
	}
	if l2==nil{
		return  l1
	}
	more:=0
	var head *ListNode
	var tail *ListNode
	for l1!=nil||l2!=nil{
		n1,n2:=0,0
		if l1!=nil{
			n1=l1.Val
			l1=l1.Next
		}
		if l2!=nil{
			n2=l2.Val
			l2=l2.Next
		}
		sum:=n1+n2+more
		sum,more=sum%10,sum/10
		if head==nil{
			head=&ListNode{Val: sum}
			tail=head
		}else {
			tail.Next=&ListNode{Val: sum}
			tail=tail.Next
		}
	}
	if more==1{
		tail.Next=&ListNode{Val:more}
	}
	return head
}