package leetcdoe

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	_, right := getNearGreaterNoRepeat(nums2)
	m := make(map[int]int, 0)
	for i, v := range nums2 {
		m[v] = i
	}
	ret := make([]int, 0, len(nums1))
	for _, i2 := range nums1 {
		if right[m[i2]] >= 0 {
			ret = append(ret, nums2[right[m[i2]]])
		} else {
			ret = append(ret, right[m[i2]])
		}
	}
	return ret
}

func getNearGreaterNoRepeat(nums []int) (left, right []int) {
	left = make([]int, len(nums))
	right = make([]int, len(nums))

	stack := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			//pop
			popIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				left[popIndex] = -1
			} else {
				left[popIndex] = stack[len(stack)-1]
			}
			right[popIndex] = i
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		popIndex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			left[popIndex] = -1
		} else {
			left[popIndex] = stack[len(stack)-1]
		}
		right[popIndex] = -1
	}
	return left, right
}
