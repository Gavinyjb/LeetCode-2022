package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	low,hight:=0,len(nums)
	for low<hight{
		mid:=low+(hight-low)>>1
		if nums[mid]==target{
			return mid
		}else if nums[mid]<target {
			low=mid+1
		}else {
			hight=mid
		}
	}
	//分别处理如下四种情况
	//目标值在数组所有元素之前 [0,0)
	//目标值等于数组中某一个元素 return mid
	//目标值插入数组中的位置 [low, hight) ，return hight 即可
	//目标值在数组所有元素之后的情况 [low, hight)，return hight 即可
	return hight
}
func main() {
	nums:=[]int{-1,0,3,5,9,12}
	target := 2
	ret:=searchInsert(nums,target)
	fmt.Println(ret)
}

abbbbab ab
abbbcab abc