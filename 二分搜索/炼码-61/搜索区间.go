package 炼码_61

func SearchRange(a []int, target int) []int {
	// write your code here
	if len(a) == 0 {
		return []int{-1, -1}
	}
	result := make([]int, 2)
	start, end := 0, len(a)-1
	for start+1 < end {
		mid := start + (end-start)>>1
		if a[mid] == target {
			//如果相等，应该继续向左找，就能找到第一个目标值的位置
			end = mid
		} else if a[mid] < target {
			start = mid
		} else if a[mid] > target {
			end = mid
		}
	}
	//搜索左边的索引
	if a[start] == target {
		result[0] = start
	} else if a[end] == target {
		result[0] = end
	} else {
		result[0] = -1
		result[1] = -1
		return result
	}
	start, end = 0, len(a)-1
	for start+1 < end {
		mid := start + (end-start)>>1
		if a[mid] == target {
			// 如果相等，应该继续向右找，就能找到最后一个目标值的位置
			start = mid
		} else if a[mid] > target {
			end = mid
		} else if a[mid] < target {
			start = mid
		}
	}
	// 搜索右边的索引
	if A[end] == target {
		result[1] = end
	} else if A[start] == target {
		result[1] = start
	} else {
		result[0] = -1
		result[1] = -1
		return result
	}
	return result
}
