package offer40

import (
	"container/heap"
	"fmt"
)

//以下代码用快排实现
func getLeastNumbers(arr []int, k int) []int {
	if k == 0 || len(arr) == 0 {
		return []int{}
	}
	if k > len(arr) {
		return arr
	}
	var quickSort func(arr []int, start, end int) []int
	quickSort = func(arr []int, start, end int) []int {
		p := partition(arr, start, end)
		if p == k-1 {
			return arr[:k]
		} else if p < k-1 {
			return quickSort(arr, p+1, end)
		} else {
			return quickSort(arr, start, p-1)
		}
	}
	return quickSort(arr, 0, len(arr)-1)
}
func partition(nums []int, start, end int) int {
	l, r := start, end
	x := nums[start]
	for l < r {
		for l < r && nums[r] > x {
			r--
		}
		for l < r && nums[l] <= x {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[l], nums[start] = nums[start], nums[l]
	fmt.Println(nums, "i:", l, "x:", x)
	return l
}

//以下代码用堆实现
type intHeap []int

func (h intHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h intHeap) Len() int {
	return len(h)
}
func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func getLeastNumbers2(arr []int, k int) []int {
	myHeap := &intHeap{}
	heap.Init(myHeap)
	for _, v := range arr {
		heap.Push(myHeap, v)
		if myHeap.Len() > k {
			heap.Pop(myHeap)
		}
	}
	ret := make([]int, 0)
	for myHeap.Len() > 0 {
		ret = append(ret, heap.Pop(myHeap).(int))
	}
	return ret
}
