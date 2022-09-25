package Problem_0084_LargestRectangleInHistogram

func largestRectangleArea(heights []int) int {
	if heights == nil || len(heights) == 0 {
		return 0
	}
	maxArea := 0

	stack := make([]int, 0) //存储下标

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] <= heights[stack[len(stack)-1]] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			k := -1
			if len(stack) == 0 {
				k = -1
			} else {
				k = stack[len(stack)-1]
			}
			curArea := (i - k - 1) * heights[j]
			maxArea = max(maxArea, curArea)
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		j := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k := -1
		if len(stack) == 0 {
			k = -1
		} else {
			k = stack[len(stack)-1]
		}
		curArea := (len(heights) - k - 1) * heights[j]
		maxArea = max(maxArea, curArea)
	}
	return maxArea
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
