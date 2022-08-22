package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } //大根堆
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

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

//3 2
//1 1
//1 2
//1 3

func main() {
	n, m := 0, 0
	fmt.Scan(&n)
	fmt.Scan(&m)
	yList := make([]int, n) //个数
	xList := make([]int, n) //单价
	for i := 0; i < n; i++ {
		fmt.Scan(&yList[i])
		fmt.Scan(&xList[i])
	}
	fmt.Println(yList)
	fmt.Println(xList)
	h := &IntHeap{}
	heap.Init(h) //初始化大根堆

	for i := 0; i < n; i++ {
		for j := 0; j < yList[i]; j++ { //每家店有几个商品就往堆里面放几个价格进去
			heap.Push(h, xList[i])
			if h.Len() > m { //数量超过m之后就从堆顶弹出最大的
				heap.Pop(h)
			}
		}
	}
	//遍历堆中所有累加得结果
	res := 0
	for h.Len() > 0 {
		res += heap.Pop(h).(int)
	}
	fmt.Println(res)
}
