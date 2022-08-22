package main

import (
	"fmt"
	"sort"
)

func main() {
	n, m := 0, 0
	fmt.Scanf("%d %d\n", &n, &m)
	aList := make([]int, n)
	bList := make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Scan(&aList[i])
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&bList[i])
	}

	sort.Ints(aList)
	sort.Ints(bList)
	//fmt.Println(aList)
	//fmt.Println(bList)

	if bList[m-1] < aList[n-1] {
		fmt.Println(-1)
	} else {
		weight := 0
		for i, j := 0, 0; i < n && j < m; i++ {
			for aList[i] > bList[j] {
				j++
			}
			weight += bList[j]
		}
		fmt.Println(weight)
	}

	//fmt.Println(-1)
}
