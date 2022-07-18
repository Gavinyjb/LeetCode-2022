package leetcdoe

import "fmt"

func singleNumber(nums []int) []int {
	// a=a^b
	// b=a^b
	// a=a^b
	// 关键点怎么把a^b分成两部分,方案：可以通过diff最后一个1区分

	diff := 0
	for i := 0; i < len(nums); i++ {
		diff ^= nums[i]
	}
	fmt.Printf("diff:%b\n", diff)
	fmt.Printf("3^5:%b\n", 3^5)
	result := []int{diff, diff}
	// 去掉末尾的1后异或diff就得到最后一个1的位置
	diff = (diff & (diff - 1)) ^ diff
	//以最低位的1为划分，将所有元素分为两类
	for i := 0; i < len(nums); i++ {
		if diff&nums[i] == 0 {
			result[0] ^= nums[i]
		} else {
			result[1] ^= nums[i]
		}
	}
	return result
}
