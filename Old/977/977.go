package main

import "sort"
//方法一
func sortedSquares1(nums []int) []int {
	for i, num := range nums {
		nums[i]=num*num
	}
	sort.Ints(nums)
	return nums
}
func sortedSquares(nums []int) []int {
	low,high,k:=0,len(nums)-1,len(nums)-1
	ans:=make([]int,len(nums))
	for low<=high{
		if nums[low]*nums[low]>nums[high]*nums[high]{
			ans [k]=nums[low]*nums[low]
			low++
		}else {
			ans [k]=nums[high]*nums[high]
			high--
		}
		k--
	}
	return ans
}

func main() {

}
