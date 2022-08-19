package problem00153sum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans :=make([][]int,0)
	// 第一个数选了i位置的数
	for i := 0; i < len(nums)-2; i++ {
		var nexts [][]int
		if i==0||nums[i-1]!=nums[i]{
			nexts=twoSum(nums,i+1,-nums[i])
			for idx:= range nexts {
				nexts[idx]=append(nexts[idx], nums[i])
			}
		}
		ans = append(ans, nexts...)
	}
	return ans

}

// nums已经有序了
// nums[begin......]范围上，找到累加和为target的所有二元组
func twoSum(nums []int, begin int, target int) [][]int {
	ret := make([][]int, 0)
	l, r := begin, len(nums)-1
	for l < r {
		if nums[l]+nums[r] < target {
			l++
		} else if nums[l]+nums[r] > target {
			r--
		} else {
			if l == begin || nums[l] != nums[l-1] { //结果去重
				ret = append(ret, []int{nums[l], nums[r]})
			}
			l++
		}
	}
	return ret
}