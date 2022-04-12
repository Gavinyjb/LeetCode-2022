package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	low:=0 // 滑动窗口起始位置
	sum := 0 // 滑动窗口数值之和
	result:=len(nums)+1// 初始化返回长度为l+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况
	for high := 0; high < len(nums); high++ {
		sum+=nums[high]
		for sum>=target{
			result=min(result,high-low+1)
			sum-=nums[low]
			low++
		}
	}
	if result==len(nums)+1{
		return 0
	}
	return result
}
func min(a,b int) int {
	if a<b{
		return a
	}else {
		return b
	}
}
func main() {
	input:=[]int{2,3,1,2,4,3}
	target:=7
	ret:=minSubArrayLen(target,input)
	fmt.Println(ret)
}
