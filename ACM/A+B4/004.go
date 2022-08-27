package main

import "fmt"

//https://ac.nowcoder.com/acm/contest/5652/D
func main() {
	n := -1
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		} else {
			nums := make([]int, n)
			sum := 0
			for i := 0; i < n; i++ {
				fmt.Scan(&nums[i])
				sum += nums[i]
			}
			fmt.Println(sum)
		}
	}
}
