package leetcode

func trap(height []int) int {
	left, right := getNearGreaterNoRepeat(height)
	res := 0
	//fmt.Println(left)
	//fmt.Println(right)
	for i := 0; i < len(height); i++ {
		list := make([]int, 0)
		for left[i] >= 0 && right[i] >= 0 {
			list = append(list, i)
			i++
		}
		for _, v := range list {
			currHeight := min(height[left[list[0]]], height[right[list[len(list)-1]]]) - height[v]
			res += currHeight * 1
		}
	}
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getNearGreaterNoRepeat(nums []int) (left, right []int) {
	left = make([]int, len(nums))
	right = make([]int, len(nums))

	//从栈顶至栈底递增（即栈顶最小） 栈中存储的是下标
	stack := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		//!!!!!!!接雨水  应为下一个大于等于的高度
		for len(stack) > 0 && nums[i] >= nums[stack[len(stack)-1]] {
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
