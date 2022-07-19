package leetcode

func searchInsert(nums []int, target int) int {
	//思路:找到第一个>= target 的元素位置
	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := start + (end-start)>>1
		if nums[mid] == target {
			//标记开始的位置
			start = mid
		} else if nums[mid] > target {
			end = mid
		} else if nums[mid] < target {
			start = mid
		}
	}
	if nums[start] >= target {
		return start
	} else if nums[end] >= target {
		return end
	} else if nums[end] < target { //目标值比所有值都大
		return end + 1
	}
	return 0
}
