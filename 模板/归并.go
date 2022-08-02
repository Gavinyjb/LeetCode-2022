package 模板

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) >> 1
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	res := merge(left, right)
	return res
}
func merge(left, right []int) []int {
	ret := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			ret = append(ret, left[i])
			i++
		} else {
			ret = append(ret, right[j])
			j++
		}
	}
	ret = append(ret, left[i]...)
	ret = append(ret, right[j]...)
	return ret
}
