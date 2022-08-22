package Problem_0042_TrappingRainWater

func trap(height []int) int {
	if height == nil || len(height) < 3 {
		return 0
	}
	l, r := 1, len(height)-2
	leftMax, rightMax := height[0], height[len(height)-1]
	ans := 0
	for l <= r {
		if leftMax <= rightMax {
			ans += max(0, leftMax-height[l])
			leftMax = max(leftMax, height[l])
			l++
		} else {
			ans += max(0, rightMax-height[r])
			rightMax = max(rightMax, height[r])
			r--
		}
	}
	return ans
}

//单调栈
func trap2(height []int) int {
	if len(height) < 3 {
		return 0
	}
	gotLeft, gotRight := getNearGreaterNoRepeat(height)
	//fmt.Println(gotLeft,gotRight)
	ans := 0
	for i := 1; i < len(height)-2; i++ {
		if gotLeft[i] != -1 && gotRight[i] != -1 {
			//fmt.Println(i)
			ans += max(0, min(height[gotRight[i]], height[gotLeft[i]])-height[i]) * (gotRight[i] - gotLeft[i] - 1)
		}
	}
	return ans
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getNearGreaterNoRepeat(nums []int) (left, right []int) {
	left = make([]int, len(nums))
	right = make([]int, len(nums))

	//从栈顶至栈底递增(即栈顶是最小) 栈中存储的小标
	stack := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			popIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			//的栈顶即为左侧第一个大于的数字的下标
			if len(stack) == 0 {
				left[popIdx] = -1
			} else {
				left[popIdx] = stack[len(stack)-1]
			}
			right[popIdx] = i
		}
		stack = append(stack, i)
	}
	//处理栈中剩余的下标元素
	for len(stack) > 0 {
		popIdx := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		//的栈顶即为左侧第一个大于的数字的下标
		if len(stack) == 0 {
			left[popIdx] = -1
		} else {
			left[popIdx] = stack[len(stack)-1]
		}
		right[popIdx] = -1
	}
	return left, right
}
