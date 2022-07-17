package leetcdoe

func dailyTemperatures(temperatures []int) []int {
	left, right := getNearGreaterNoRepeat(temperatures)
	result := make([]int, 0, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		result = append(result, max(0, right[i]-left[i]))
	}
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//向左向右遇到的第一个比自己大的数字的下标
/*
元数据：[73,74,75,71,69,72,76,73]
left： [-1 -1 -1 2 3 2 -1 6]
right：[1 2 6 5 5 6 -1 -1]
*/
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
