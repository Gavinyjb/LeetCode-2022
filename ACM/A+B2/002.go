package main

import "fmt"

//https://ac.nowcoder.com/acm/contest/5652/B

func main() {
	var t int
	fmt.Scan(&t)
	var a, b int
	for i := 0; i < t; i++ {
		fmt.Scan(&a, &b)
		fmt.Println(a + b)
	}
}
