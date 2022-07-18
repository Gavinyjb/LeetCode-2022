package leetcode

func nextGreaterElements(nums []int) []int {
	_, right := getNearGreaterNoRepeat(nums)
	ret := make([]int, 0, len(nums))
	for _, v := range right {
		if v >= 0 {
			ret = append(ret, nums[v])
		} else {
			ret = append(ret, v)
		}
	}
	return ret
}

func getNearGreaterNoRepeat(nums []int) (left, right []int) {
	left = make([]int, len(nums))
	right = make([]int, len(nums))
	for i := range right {
		right[i] = -1
	}
	//单调递减 存储数组下标索引
	stack := make([]int, 0, len(nums))
	l := len(nums)
	for i := 0; i < 2*l; i++ {
		for len(stack) > 0 && nums[i%len(nums)] > nums[stack[len(stack)-1]] {
			//pop
			popIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				left[popIndex] = -1
			} else {
				left[popIndex] = stack[len(stack)-1]
			}
			right[popIndex] = i % l
		}
		stack = append(stack, i%l)
	}
	for len(stack) > 0 {
		popIndex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			left[popIndex] = -1
		} else {
			left[popIndex] = stack[len(stack)-1]
		}
		//right[popIndex] = -1
	}
	return left, right
}
