package main

import "fmt"

func f(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n%2 == 0 { //偶数
		return n - 1 + f(n/2) + f(n/2-1)
	} else {
		return n - 1 + 2*f(n/2)
	}
}
func process(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	dp[2] = 1
	for i := 3; i <= n; i++ {
		if i%2 == 0 {
			dp[i] = i - 1 + dp[i/2] + dp[i/2-1]
		} else {
			dp[i] = i - 1 + 2*dp[i/2]
		}
	}
	return dp[n]
}
func main() {
	n := 5
	res := process(n) + n //需要加n
	fmt.Println(res)
}
