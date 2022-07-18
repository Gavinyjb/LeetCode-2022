package leetcode

import "fmt"

func singleNumber(nums []int) int {
	// 统计每位1的个数
	var result int
	for i := 0; i < 64; i++ {
		sum := 0
		for j := 0; j < len(nums); j++ {
			// 统计1的个数
			sum += (nums[j] >> i) & 1
		}
		fmt.Printf("sum:%v\n", sum)
		// 还原位00^10=10 或者用| 也可以
		result ^= (sum % 3) << i
	}
	return result
}
