package Problem_0033_SearchInRotatedSortedArray

//nums 原本是有序数组
//结论: 只要它不都一样，它总能二分
//要点: 能二分就尽可能二分
func search(nums []int, target int) int {
	//思路：/ /两条上升直线，四种情况判断
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	mid := 0
	for l+1 < r {
		mid = l + (r-l)/2
		//相等直接返回
		if nums[mid] == target {
			return mid
		}
		//判断在哪个区间，可能分为四种情况
		if nums[l] < nums[mid] { //mid在左边的线段
			if nums[l] <= target && nums[mid] >= target {
				r = mid
			} else {
				l = mid
			}
		} else if nums[r] > nums[mid] { //mid在右边的线段
			if nums[mid] <= target && nums[r] >= target {
				l = mid
			} else {
				r = mid
			}
		}
	}
	if nums[l] == target {
		return l
	} else if nums[r] == target {
		return r
	} else {
		return -1
	}
}
