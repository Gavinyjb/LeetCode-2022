package main

import "fmt"

func help(x1, y1, x2, y2 int) int {
	m := abs(y2, y1) + abs(x2, x1)
	//fmt.Printf("m=%v\n", m)
	n := abs(x2, x1)
	//fmt.Printf("n=%v\n", n)
	ms := jeicheng(m)
	//fmt.Printf("ms=%v\n", ms)
	ns := jeicheng(n)
	//fmt.Printf("ns=%v\n", ns)
	return ms / ns
}
func jeicheng(m int) int {
	ret := 1
	for i := 1; i <= m; i++ {
		ret = i * ret
	}
	return ret
}
func abs(a, b int) int {
	if a-b > 0 {
		return a - b
	} else {
		return b - a
	}
}
func main() {
	n := 0
	fmt.Scanf("%v\n", &n)
	x1, y1 := 0, 0
	fmt.Scanf("%v %v\n", &x1, &y1)
	x2, y2 := 0, 0
	fmt.Scanf("%v %v\n", &x2, &y2)
	//fmt.Println(x2)
	//fmt.Println(y2)
	numAll := help(1, 1, n, n)
	num1ton := help(x1, y1, n, n)
	num2ton := help(x2, y2, n, n)
	num0to1 := help(1, 1, x1, y1)
	num0to2 := help(1, 1, x2, y2)
	if x1 == x2 && y1 == y2 {
		res := numAll - num0to1*num1ton
		fmt.Println(res % 998244353)
	} else {
		res := numAll - num0to2*num2ton - num0to1*num1ton
		fmt.Println(res % 998244353)
	}
}
