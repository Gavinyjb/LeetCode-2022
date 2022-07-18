package leetcode

import "fmt"

func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		num = num & (num - 1)
		res++
	}
	return res
}
func HammingWeight(n uint32) uint32 {
	fmt.Printf("%032b\n", n)
	n = (n & 0x55555555) + ((n >> 1) & 0x55555555)
	fmt.Printf("%032b\n", n)
	n = (n & 0x33333333) + ((n >> 2) & 0x33333333)
	fmt.Printf("%032b\n", n)
	n = (n & 0x0f0f0f0f) + ((n >> 4) & 0x0f0f0f0f)
	fmt.Printf("%032b\n", n)
	n = (n & 0x00ff00ff) + ((n >> 8) & 0x00ff00ff)
	fmt.Printf("%032b\n", n)
	n = (n & 0x0000ffff) + ((n >> 16) & 0x0000ffff)
	fmt.Printf("%032b\n", n)
	return n
}
func main() {
	var n uint32
	n = 9
	HammingWeight(n)
}
