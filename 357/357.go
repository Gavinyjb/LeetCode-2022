package main

import (
	"fmt"
	"strconv"
)
//解法一 暴力！超时
//func countNumbersWithUniqueDigits(n int) int {
//	max:=1
//	for i := 0; i < n; i++ {
//		max*=10
//	}
//	ret:=0 //返回值
//	//fmt.Println(max)
//	for i := 0; i < max; i++ {
//		if hasDup(i){
//			ret++
//		}
//	}
//	return ret
//}
//有重复数返回false否则true
func hasDup(num int)bool  {
	hashMap:=make(map[string]int,0)
	str:=strconv.Itoa(num)
	for len(str)>0{
		key:=str[:1]
		str=str[1:]
		if hashMap[key]>0{
			return false
		}
		hashMap[key]++
	}
	return true
}
func countNumbersWithUniqueDigits(n int) int {
	ret:=1
	if n==0{
		return ret
	}
	for i := 1; i <= n; i++ {
		ret+=calculate(i,10)-calculate(i-1,9)
	}
	return ret
}
//计算全排列A(m,n)  m<=n
func calculate(m,n int) int {
	ret:=1
	for i := 0; i <m ; i++ {
		ret=ret*n
		n--
	}
	return ret
}

func main() {
	fmt.Println(calculate(0,0))
	fmt.Println(countNumbersWithUniqueDigits(2))
	//ret:=hasDup(0)
	//ret2:=countNumbersWithUniqueDigits(2)
	//fmt.Println(ret)
	//fmt.Println(ret2)
}
