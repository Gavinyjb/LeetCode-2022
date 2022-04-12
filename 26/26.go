package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums)<2{
		return len(nums)
	}
	slow,fast:=1,1
	for ;fast<len(nums);fast++{
		if nums[fast]!=nums[fast-1]{
			nums[slow]=nums[fast]
			slow++
		}
	}
	fmt.Println(nums)
	return slow
}
func main() {
	input:=[]int{2,2,3,3,4,5,5,5,6,6,6,6,6,}
	//val:=3
	ret:=removeDuplicates(input)
	fmt.Println(ret)
}
