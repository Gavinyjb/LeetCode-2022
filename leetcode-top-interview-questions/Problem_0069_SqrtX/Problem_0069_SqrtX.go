package Problem_0069_SqrtX

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x < 3 {
		return 1
	}
	ans := 1
	l, r, mid := 1, x, 0
	for l <= r {
		mid = (l + r) / 2
		if mid*mid <= x {
			l = mid + 1
			ans = mid
		} else {
			r = mid - 1
		}
	}
	return ans
}
