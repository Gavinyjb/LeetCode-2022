package 模板

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
