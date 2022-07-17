package main

import "fmt"

func solution(x, y, a, b int) int {
	if b%a != 0 {
		return -1
	}
	devide := b / a
	mul := x * y
	ans := 0
	if devide%mul == 0 {
		ans += devide / mul * 2
		devide /= mul
	}
	if devide%y != 0 && devide%x != 0 {
		return -1
	}
	if devide%y == 0 {
		ans += devide / y
	} else {
		ans += devide / x
	}
	return ans
}
func main() {
	x, y, a, b := 0, 0, 0, 0
	fmt.Scan(&x, &y, &a, &b)
	res := solution(x, y, a, b)
	fmt.Println(res)
}
