package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	//虚拟一个头结点，使得 头结点操作与其他节点操作统一
	newHead := &ListNode{Val: 0,Next: head}
	pre:=newHead
	cur:=head
	for cur!=nil{
		if cur.Val==val{
			pre.Next=cur.Next
		}else {
			pre=cur
		}
		cur=cur.Next
	}
	return newHead.Next
}
func main() {

}
