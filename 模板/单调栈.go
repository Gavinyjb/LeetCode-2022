package 模板

//向左向右遇到的第一个比自己大的数字的下标
/*
元数据：[73,74,75,71,69,72,76,73]
left： [-1 -1 -1 2 3 2 -1 6]
right：[1 2 6 5 5 6 -1 -1]
*/
func getNearGreaterNoRepeat(nums []int) (left, right []int) {
	left = make([]int, len(nums))
	right = make([]int, len(nums))

	//从栈顶至栈底递增（即栈顶最小） 栈中存储的是下标
	stack := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			popIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			//新的栈顶即为左侧第一个大于的数字的下标
			if len(stack) == 0 {
				left[popIndex] = -1
			} else {
				left[popIndex] = stack[len(stack)-1]
			}
			right[popIndex] = i
		}
		stack = append(stack, i)
	}
	//处理栈中剩余的下标元素
	for len(stack) > 0 {
		popIndex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		//新的栈顶即为左侧第一个大于的数字的下标
		if len(stack) == 0 {
			left[popIndex] = -1
		} else {
			left[popIndex] = stack[len(stack)-1]
		}
		right[popIndex] = -1
	}
	return left, right
}

//向左向右遇到的第一个比自己小的数字的下标
func getNearLessNoRepeat(nums []int) (left, right []int) {
	left = make([]int, len(nums))
	right = make([]int, len(nums))

	//从栈顶至栈底递减（即栈顶最大） 栈中存储的是下标
	stack := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[i] < nums[stack[len(stack)-1]] {
			popIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			//新的栈顶即为左侧第一个大于的数字的下标
			if len(stack) == 0 {
				left[popIndex] = -1
			} else {
				left[popIndex] = stack[len(stack)-1]
			}
			right[popIndex] = i
		}
		stack = append(stack, i)
	}
	//处理栈中剩余的下标元素
	for len(stack) > 0 {
		popIndex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		//新的栈顶即为左侧第一个大于的数字的下标
		if len(stack) == 0 {
			left[popIndex] = -1
		} else {
			left[popIndex] = stack[len(stack)-1]
		}
		right[popIndex] = -1
	}
	return left, right
}
