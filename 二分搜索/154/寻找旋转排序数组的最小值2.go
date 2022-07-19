package leetcode

func findMin(nums []int) int {
	//思路:跳过重复元素,mid 值和end比较，分两种情况进行处理
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1
	for start+1 < end {
		//去除重复元素
		for start < end && nums[end] == nums[end-1] {
			end--
		}
		for start < end && nums[start] == nums[start+1] {
			start++
		}
		mid := start + (end-start)>>1
		//中间元素和最后一个元素比较(判断中间点落在左侧上升区，还是右侧上升去
		if nums[mid] <= nums[end] {
			end = mid
		} else {
			start = mid
		}
	}
	if nums[start] > nums[end] {
		return nums[end]
	}
	return nums[start]
}
