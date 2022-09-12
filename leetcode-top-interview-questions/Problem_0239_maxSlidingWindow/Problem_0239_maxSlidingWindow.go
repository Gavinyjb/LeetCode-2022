package Problem_0239_maxSlidingWindow

import (
	"container/heap"
	"fmt"
)

type intHeap [][]int //0: idx 1：num

func (h intHeap) Less(i, j int) bool {
	return h[i][1] > h[j][1]
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
	item := old[n-1]
	*h = old[:n-1]
	return item
}
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}
func maxSlidingWindow(nums []int, k int) []int {
	ret := make([]int, 0)
	var h intHeap
	heap.Init(&h)
	for i := 0; i < k-1; i++ {
		num := nums[i]
		one := []int{i, num}
		fmt.Println(h)
		heap.Push(&h, one)
	}
	for i := k - 1; i < len(nums); i++ {
		num := nums[i]
		one := []int{i, num}
		fmt.Println(h)
		heap.Push(&h, one)

		leave := heap.Pop(&h).([]int)
		for leave[0] <= i-k {
			fmt.Println(h)
			leave = heap.Pop(&h).([]int)
		}
		heap.Push(&h, leave)
		ret = append(ret, leave[1])
	}
	return ret
}

//单调栈
func maxSlidingWindow1(nums []int, k int) []int {
	ret := make([]int, 0)
	stack := make([]int, 0) //存储下标  自左向右单调递减

	for i, num := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= num {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		if i-k == stack[0] {
			stack = stack[1:]
		}
		if i >= k-1 {
			ret = append(ret, nums[stack[0]])
		}
	}
	return ret
}
