package main

import (
	"fmt"
	"sort"
)

//第一行两个正整数n和m，表示总试题数和最多复习试题数。
//
//接下来一行n个整数，分别为p1 p2...pn，表示小美有pi%的概率，即pi=pi/100的概率做对第i个题。（注意，这里为了简单起见，将概率pi扩张100倍成为整数pi方便输入）
//
//接下来一行n个整数，分别表示a1 a2...an，分别表示第 i 个题做对的分值。
//
//数字间两两有空格隔开，对于所有数据，1≤m≤n≤50000,0≤pi≤100,1≤ai≤1000

//输出一行一个恰好两位的小数，表示能获得的最大期望总分。（如果答案为10应输出10.00，2.5应输出2.50）

//
//2 1
//89 38
//445 754

//1150.05

//如果都不复习，小美总分的期望为89%*445+38%*754=682.57
//
//如果复习第一道题，小美总分的期望为100%*445+38%*754=731.52
//
//如果复习第二道题，小美总分的期望为89%*445+100%*754=1150.05
//
//所以选择复习第二道题，这样能获得最大期望总分1150.05
//
//根据每题复习后的收益进行排序即可
//rest 剩余rest 可以复习

func main() {
	n, m := 0, 0
	fmt.Scanf("%d %d\n", &n, &m)

	pList := make([]int, n)
	aList := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&pList[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&aList[i])
	}
	fmt.Println(pList)
	fmt.Println(aList)
	//如果可以获得100分可以多提升多少分
	upList := make([]int, n)
	total := 0 //总分
	for i := 0; i < n; i++ {
		total += pList[i] * aList[i]
		upList[i] = (100 - pList[i]) * aList[i]
	}
	fmt.Println(upList)
	sort.Ints(upList)

	for i := 0; i < m; i++ {
		total += upList[n-i-1]
	}
	result := float64(total) / 100.0
	fmt.Printf("%.2f\n", result)

}
