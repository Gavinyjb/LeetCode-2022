package leetcode_78
//package main

import "fmt"

func subsets(nums []int) [][]int {
	res:=make([][]int,0)
	track:=make([]int,0)
	var backtrack func(nums []int,start int)
	backtrack=func(nums []int,start int)  {
		//触发结束条件
		//前序位置，每个节点的值都是一个子集
		temp:=make([]int,len(track))
		copy(temp,track)
		res = append(res, temp)
		//for 选择 in 选择列表:
		for i := start; i < len(nums); i++ {
			//做选择
			track=append(track,nums[i])
			backtrack(nums,i+1)
			track=track[:len(track)-1]
		}
	}
	backtrack(nums,0)
	return res
}
func main()  {
	input:=[]int{1,2,3}
	ret:=subsets(input)
	fmt.Println(ret)
}

