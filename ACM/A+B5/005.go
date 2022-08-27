package main

import "fmt"

//https://ac.nowcoder.com/acm/contest/5652/E
func main() {
	t := 0
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		n := 0
		fmt.Scan(&n)
		nums := make([]int, n)
		sum := 0
		for k := 0; k < n; k++ {
			fmt.Scan(&nums[k])
			sum += nums[k]
		}
		fmt.Println(sum)
	}
}
