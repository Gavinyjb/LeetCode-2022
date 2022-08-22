package main

import (
	"fmt"
)

//第一行一个正整数n，表示棋盘大小。
//
//第二行两个整数，分别表示x1与y1，即第一个定位器的位置。
//
//第三行两个整数，分别表示x2与y2，即第二个定位器的位置。
//
//第四行两个整数，分别表示x3与y3，即第三个定位器的位置。
//
//第五行三个整数，分别表示第一、二、三个定位器到信标的曼哈顿距离。第i个定位器到信标的曼哈顿距离即(|xi-a|+|yi-b|)
//
//数字间两两有空格隔开，对于所有数据， n≤50000,  1≤xi,yi≤n

//输出一行两个整数，表示字典序最小的可能的信标位置。
//3
//2 1
//2 2
//2 3
//2 1 2
////out
//1 2
//type indexTwo struct {
//	x int
//	y int
//}
//
//func (t *indexTwo) String() string {
//	ret:=""
//	ret+=strconv.Itoa(t.x)
//	ret+=strconv.Itoa(t.y)
//	return ret
//}
func abs(a, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}
func main() {
	n := 0
	fmt.Scan(&n)
	x1, x2, x3, y1, y2, y3 := -1, -1, -1, -1, -1, -1
	fmt.Scan(&x1)
	fmt.Scan(&y1)
	fmt.Scan(&x2)
	fmt.Scan(&y2)
	fmt.Scan(&x3)
	fmt.Scan(&y3)

	distList := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scan(&distList[i])
	}
	//fmt.Println(x1,x2,x3,y1,y2,y3)
	result := make([]int, 2)
	flag := false
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if abs(i, x1)+abs(j, y1) == distList[0] && abs(i, x2)+abs(j, y2) == distList[1] && abs(i, x3)+abs(j, y3) == distList[2] {
				result[0] = i
				result[1] = j
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	if flag {
		fmt.Printf("%v %v\n", result[0], result[1])
	} else {
		result[0], result[1] = -1, -1
		fmt.Printf("%v %v\n", result[0], result[1])
	}
}
