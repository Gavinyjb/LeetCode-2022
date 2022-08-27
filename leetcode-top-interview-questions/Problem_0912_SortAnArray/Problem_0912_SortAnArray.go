package Problem_0912_SortAnArray

func quickSortMine(nums []int, start, end int) {
	if start < end {
		l, r := twoPartition(nums, start, end)
		quickSortMine(nums, start, l-1)
		quickSortMine(nums, r+1, end)
	}
}
func twoPartition(nums []int, start, end int) (less, more int) {
	if start > end {
		return -1, -1
	}
	if start == end {
		return start, end
	}
	less, more = start-1, end+1
	cur := start
	x := nums[end]
	//all in [start,less]  < x
	//all in [more,end]    > x
	//all in (less,cur)    =x

	//终止条件  cur==more
	for cur < more {
		if nums[cur] < x {
			less++
			swap(nums, cur, less)
			cur++
		} else if nums[cur] == x {
			cur++
		} else {
			more--
			swap(nums, more, cur)
		}
	}
	return less + 1, more - 1
}
func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

//经典快排
func quickSort(nums []int, start, end int) {
	if start < end {
		p := partition(nums, start, end)
		quickSort(nums, start, p-1)
		quickSort(nums, p+1, end)
	}
}
func partition(nums []int, start, end int) int {
	l, r := start, end
	x := nums[start]
	for l < r {
		for l < r && nums[r] > x {
			r--
		}
		for l < r && nums[l] <= x {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[l], nums[start] = nums[start], nums[l]
	return l
}
