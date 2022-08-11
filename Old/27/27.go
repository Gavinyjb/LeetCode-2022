package main

import "fmt"

//里数组的删除并不是真的删除，只是将删除的元素移动到数组后面的空间内，然后返回数组实际剩余的元素个数，OJ 最终判断题目的时候会读取数组剩余个数的元素进行输出。
func removeElement(nums []int, val int) int {
	if len(nums)==0{
		return 0
	}
	slow,fast:=0,0
	for ;fast<len(nums);fast++{
		if nums[fast]!=val{
			if slow!=fast{
				nums[slow],nums[fast]=nums[fast],nums[slow]
			}
			slow++
		}
	}
	return slow
}

func main() {
	input:=[]int{3,2,2,3}
	val:=3
	ret:=removeElement(input,val)
	fmt.Println(ret)
}
