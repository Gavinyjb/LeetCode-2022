package problem0023mergeksortedlists

import (
	"container/heap"

	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

type ListNode struct {
	Val  int
	Next *ListNode
}
type nodeHeap []*ListNode

func (h nodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}
func (h nodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h nodeHeap) Len() int {
	return len(h)
}

func (h *nodeHeap)Pop()interface{}{
	old:=*h
	x:=old[len(old)-1]
	*h=old[:len(old)-1]
	return x
}
func (h *nodeHeap)Push(x interface{})  {
	*h=append(*h, x.(*ListNode))
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists==nil||len(lists)==0{
		return nil
	}
	var myheap nodeHeap
	heap.Init(&myheap)
	for _, node := range lists {
		if node!=nil{    //  预防这种  [[]]
			heap.Push(&myheap,node)
		}
	}
	dummy :=new(ListNode)
	cur:=dummy
	for myheap.Len()>0{
		litter:=heap.Pop(&myheap).(*ListNode)
		cur.Next=litter
		cur=cur.Next
		if litter.Next!=nil{
			heap.Push(&myheap,litter.Next)
		}
	}
	if dummy.Next==nil{
		return nil
	}else {
		return dummy.Next
	}
	
}