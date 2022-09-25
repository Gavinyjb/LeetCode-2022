package Problem_0066_PlusOne

func plusOne(digits []int) []int {
	n := len(digits)
	flag := 1
	for i := n - 1; i >= 0; i-- {
		cur := digits[i]
		if cur+flag < 10 {
			digits[i] = cur + flag
			return digits
		} else {
			digits[i] = 0
		}
	}
	ans := make([]int, 1)
	ans[0] = 1
	ans = append(ans, digits...)
	return ans
}
