package leetcode

func rangeBitwiseAnd(m int, n int) int {
	// m 5 1 0 1
	//   6 1 1 0
	// n 7 1 1 1
	// 把可能包含0的全部右移变成
	// m 5 1 0 0
	//   6 1 0 0
	// n 7 1 0 0
	// 所以最后结果就是m<<count
	var count int
	for m != n {
		m >>= 1
		n >>= 1
		count++
	}
	return m << count
}
