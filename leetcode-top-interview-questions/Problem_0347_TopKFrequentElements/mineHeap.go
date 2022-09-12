package Problem_0347_TopKFrequentElements

import "container/heap"

//0:nums 1:count
type numsHeap [][2]int

func (h numsHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}
func (h numsHeap) Len() int {
	return len(h)
}
func (h numsHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *numsHeap) Push(item interface{}) {
	*h = append(*h, item.([2]int))
}
func (h *numsHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	myHeap := &numsHeap{}
	for num, count := range m {
		item := [2]int{num, count}
		heap.Push(myHeap, item)
		if myHeap.Len() > k {
			heap.Pop(myHeap)
		}
	}
	ret := make([]int, myHeap.Len())
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(myHeap).([2]int)[0]
	}
	return ret
}
