package Problem_0029_DivideTwoIntegers

import (
	"math"
	"strings"
)

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

//除法  返回商和余数
func div(a, b int) (int, int) { //保证a,b 都不是系统最小

	dividend, divisor := 0, 0 //被除数 除数
	if isNeg(a) {
		dividend = negNum(a)
	} else {
		dividend = a
	}
	if isNeg(b) {
		divisor = negNum(b)
	} else {
		divisor = b
	}
	quotient := 0              //商
	remainder := 0             //余数
	for i := 31; i >= 0; i-- { //i:=31;i>-1;i--  不可以用+ — * \
		//比较dividend是否大于divisor的(1<<i)次方，不要将dividend与(divisor<<i)比较，而是用(dividend>>i)与divisor比较，
		//效果一样，但是可以避免因(divisor<<i)操作可能导致的溢出，如果溢出则会可能dividend本身小于divisor，但是溢出导致dividend大于divisor
		if (dividend >> i) >= divisor {
			quotient = add(quotient, 1<<i)
			dividend = minus(dividend, divisor<<i)
		}
	}
	// 确定商的符号
	if a^b < 0 {
		// 如果除数和被除数异号，则商为负数
		quotient = negNum(quotient)
	}
	//确定余数符号
	if b > 0 {
		remainder = dividend
	} else {
		remainder = negNum(dividend)
	}
	return quotient, remainder
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
	quotient, _ := div(dividend, divisor)
	return quotient
}
func help(num int) string {
	build := new(strings.Builder)
	for i := 31; i >= 0; i-- {
		if (num>>i)&1 == 0 {
			build.WriteByte('0')
		} else {
			build.WriteByte('1')
		}

	}
	return build.String()
}
