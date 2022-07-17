// 本题为考试多行输入输出规范示例，无需提交，不计分。
package main

import (
	"fmt"
)

func getMix(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func main() {
	n, m := 0, 0

	fmt.Scan(&n, &m)
	list := make([][]int, 0)
	for i := 0; i < n; i++ {
		temp := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&temp[j])
		}
		list = append(list, temp)
	}
	//fmt.Println(list)
	min := getMix(n, m)
	ret := -1
	for index := 0; index < min; index++ {
		list = list[index:][index:]
		res := -1
		for i := 1; i < min; i++ {
			temp := list[0][0] + list[i][i] + list[0][i] + list[i][0]
			if temp > res {
				res = temp
			}
		}
		if res > ret {
			ret = res
		}
	}

	fmt.Println(ret)
}
