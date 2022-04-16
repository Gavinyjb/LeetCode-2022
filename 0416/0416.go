package main

import (
	"flag"
	"fmt"
)

func main() {

	/*
	 * 数组长度为n
	 * 该数组最大值不超过k
	 * 该数组所有数都 "不同"
	 * 数组所有数之和等于x
	 */
	n := flag.Int("n", 3, "数组长度为n")
	k := flag.Int("k", 3, "该数组最大值不超过k")
	x := flag.Int("x", 6, "数组所有数之和等于x")
	flag.Parse()

	ret, isHas := solution(*n, *k, *x)
	fmt.Println(ret)
	fmt.Println(isHas)

}
func solution(n, k, x int) ([]int, bool) {
	if n == 1 {
		return []int{x}, true
	}
	/*
		情况一
	*/
	if sum(1, n) > x {
		return nil, false
	}
	/*
		情况二
	*/
	if sum(k-n+1, n) < x {
		return nil, false
	}
	/*
		情况三
	*/
	if sum(1, n) <= x && x <= sum(1, n-1)+k {
		ret := make([]int, n)
		ret[n-1] = x - sum(1, n-1)
		for i := 0; i < n-1; i++ {
			ret[i] = i + 1
		}
		return ret, true
	}
	/*
		情况四
	*/
	if sum(1, n-1)+k < x {
		ret := make([]int, n)
		ret[n-1] = k
		//subRet:=make([]int,n-1)
		subRet, isHas := solution(n-1, k-1, x-k)
		copy(ret[:n-1], subRet)
		return ret, isHas
	}

	return nil, false
}

func sum(start, num int) int {
	return (start + start + num - 1) * num / 2
}
