package main

import (
	"fmt"
	"io"
)

//https://ac.nowcoder.com/acm/contest/5652/A

func main() {

	a, b := 0, 0
	for { //循环获取输入
		_, err := fmt.Scan(&a, &b)
		if err != io.EOF {
			fmt.Println(a + b)
		} else {
			break
		}
	}
}
