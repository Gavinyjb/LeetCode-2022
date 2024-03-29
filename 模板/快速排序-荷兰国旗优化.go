package 模板

//https://leetcode.cn/problems/sort-an-array/

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

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
