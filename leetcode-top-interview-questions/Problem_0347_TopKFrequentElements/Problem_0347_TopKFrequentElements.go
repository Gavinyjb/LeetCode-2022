package Problem_0347_TopKFrequentElements

import (
	"container/heap"
	"fmt"
)

// An Item is something we manage in a priority queue.
type Item struct {
	num   int // The value of the item; arbitrary.
	count int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].count < pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
func topKFrequent(nums []int, k int) []int {
	//建立词频表
	numsMap := make(map[int]int)
	for _, num := range nums {
		numsMap[num]++
	}
	fmt.Println("map:", numsMap)
	myHeap := &PriorityQueue{}
	heap.Init(myHeap)
	i := 0
	for num, count := range numsMap {
		item := &Item{
			num:   num,
			count: count,
			index: i,
		}
		i++
		heap.Push(myHeap, item)
		showHeap(*myHeap)
		if myHeap.Len() > k {
			heap.Pop(myHeap)
		}
	}
	ret := make([]int, 0)
	for myHeap.Len() > 0 {
		item := heap.Pop(myHeap).(*Item)
		ret = append(ret, item.num)
	}
	return ret
}

func showHeap(heap PriorityQueue) {
	for i, e := range heap {
		fmt.Println(i, " ", "nums:", e.num, "count", e.count, e.index)
	}
	fmt.Println()
}
