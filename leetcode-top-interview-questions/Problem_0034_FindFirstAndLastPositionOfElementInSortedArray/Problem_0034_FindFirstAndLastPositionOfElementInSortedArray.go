package Problem_0034_FindFirstAndLastPositionOfElementInSortedArray

func searchRange(nums []int, target int) []int {
	ret := make([]int, 2)
	ret[0] = findFirst(nums, target)
	ret[1] = findLast(nums, target)
	return ret
}
func findFirst(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	mid := 0
	for l+1 < r {
		mid = l + (r-l)>>1
		if nums[mid] == target {
			r = mid
		} else if nums[mid] < target {
			l = mid
		} else if nums[mid] > target {
			r = mid
		}
	}
	//fmt.Printf("l:%d r:%d\n",l,r)
	if nums[l] == target {
		return l
	} else if nums[r] == target {
		return r
	}
	return -1
}
func findLast(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	mid := 0
	for l+1 < r {
		mid = l + (r-l)>>1
		if nums[mid] == target {
			l = mid
		} else if nums[mid] < target {
			l = mid
		} else if nums[mid] > target {
			r = mid
		}
	}
	//fmt.Printf("l:%d r:%d\n",l,r)

	if nums[r] == target {
		return r
	} else if nums[l] == target {
		return l
	}
	return -1
}
