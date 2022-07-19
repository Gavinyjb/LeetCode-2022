package leetcode

func search(nums []int, target int) int {
	//思路：/ /两条上升直线，四种情况判断
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := start + (end-start)>>1
		//相等直接返回
		if nums[mid] == target {
			return mid
		}
		//判断在哪个区间，可能分为四种情况
		if nums[start] < nums[mid] {
			if nums[start] <= target && target <= nums[mid] {
				end = mid
			} else {
				start = mid
			}
		} else if nums[end] > nums[mid] {
			if nums[end] >= target && nums[mid] <= target {
				start = mid
			} else {
				end = mid
			}
		}
	}
	if nums[start] == target {
		return start
	} else if nums[end] == target {
		return end
	}
	return -1
}
