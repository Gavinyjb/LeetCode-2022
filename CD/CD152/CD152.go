package CD152

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	//return h[i] < h[j] //小根堆
	return h[i] > h[j] //大根堆
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func FindBestKElements(nums []int, k int) []int {
	h := new(IntHeap)
	heap.Init(h)
	for _, v := range nums {
		heap.Push(h, v)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, 0)
	for h.Len() != 0 {
		ret = append(ret, heap.Pop(h).(int))
	}
	return ret
}

func main() {
	n, k := 0, 0
	fmt.Scanf("%d %d\n", &n, &k)
	//fmt.Printf("%v %v\n", n, k)
	numList := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &numList[i])
	}
	res := FindBestKElements(numList, k)
	for _, v := range res {
		fmt.Printf("%v ", v)
	}

}
