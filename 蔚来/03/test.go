// 本题为考试多行输入输出规范示例，无需提交，不计分。
package main

import (
	"fmt"
)

func main() {
	n, k := 0, 0

	fmt.Scan(&n, &k)
	list := make([][]int, 0)
	for i := 0; i < n; i++ {
		temp := make([]int, 2)
		fmt.Scan(&temp[0], &temp[1])
		list = append(list, temp)
	}
	res := 0
	for _, ints := range list {
		if ints[1] <= k {
			res += ints[0]
		}
	}
}
