package main

import "fmt"

func numberOfLines(widths []int, s string) []int {
	sum := 0
	ret1, ret2 := 1, 0
	for i := 0; i < len(s); i++ {
		temp := s[i]
		sum = sum + widths[temp-'a']
		ret2 = sum
		if sum > 100 {
			ret1++
			sum = widths[temp-'a']
			ret2 = sum
		}
	}
	return []int{ret1, ret2}
}

func main() {
	widths := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	s := "abcdefghijklmnopqrstuvwxyz"
	widths2 := []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	s2 := "bbbcccdddaaa"
	ret := numberOfLines(widths, s)
	fmt.Println(ret)
	ret2 := numberOfLines(widths2, s2)
	fmt.Println(ret2)
}
