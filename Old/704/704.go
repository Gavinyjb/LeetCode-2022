package main

import "fmt"

func search(nums []int, target int) int {
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
	return -1
}

func main() {
	nums:=[]int{-1,0,3,5,9,12}
	target := 2
	ret:=search(nums,target)
	fmt.Println(ret)
}
