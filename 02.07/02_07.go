package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	/*
		c为共有的部分（如果有）
		A:a+c
		B:b+c
	*/
	if headA == nil || headB == nil {
		return nil
	}

}
func main() {

}
