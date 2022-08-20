package Problem_0029_DivideTwoIntegers

import "math"

//加法
func add(a, b int) int {
	sum := 0
	for b != 0 {
		sum = a ^ b      //不考虑进位的情况下，
		b = (a & b) << 1 //只考虑进位的情况
		a = sum
	}
	return sum
}

//求相反数
func negNum(n int) int {
	return add(^n, 1) //取反加一
}

//减法
func minus(a, b int) int {
	return add(a, negNum(b))
}

//乘法
func multi(a, b int) int {
	res := 0
	for b != 0 {
		if (b & 1) != 0 {
			res = add(res, a)
		}
		a = a << 1
		b = b >> 1
	}
	return res
}
func isNeg(n int) bool {
	return n < 0
}

//除法
func div(a, b int) int { //保证a,b 都不是系统最小
	x, y := 0, 0
	if isNeg(a) {
		x = negNum(a)
	}
	if isNeg(b) {
		y = negNum(b)
	}
	res := 0
	for i := 31; i > negNum(1); i = minus(i, 1) { //i:=31;i>-1;i--  不可以用+ — * \
		if (x >> i) >= y {
			res = res | (1 << i)
			x = minus(x, y<<i)
		}
	}
	if (isNeg(a) == false && isNeg(b) == false) || (isNeg(a) == true && isNeg(b) == true) {
		return res
	} else {
		return negNum(res)
	}
}
func divide(dividend, divisor int) int {
	if dividend == math.MinInt32 { // 考虑被除数为最小值的情况
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == math.MinInt32 { // 考虑除数为最小值的情况
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	if dividend == 0 { // 考虑被除数为 0 的情况
		return 0
	}
	return div(dividend, divisor)
}
