package main

import (
	"fmt"
)

//func permute1(nums []int) [][]int {
//	if len(nums)<2{
//		return [][]int{nums}
//	}
//	res:=make([][]int,0)
//	for i, _ := range nums {
//		subNums:=make([]int,len(nums)-1)
//		copy(subNums,nums[:i])
//		copy(subNums[i:],nums[i+1:])
//		subres:=permute1(subNums)
//		for _, subre := range subres {
//			subre = append(subre, nums[i])
//			res = append(res, subre)
//		}
//	}
//	return res
//}
func permute(nums []int) [][]int {
	res:=make([][]int,0)
	track:=make([]int,0)
	used:=make([]bool,len(nums))
	var backtrack func(nums []int,track []int)
	backtrack= func(nums []int,track []int) {
		//触发结束条件
		if len(track)==len(nums){
			temp:=make([]int,len(track))
			copy(temp,track)
			res = append(res, temp)
		}
		//for 选择 in 选择列表:
		for i := 0; i < len(nums); i++ {
			//做选择
			if used[i]{
				continue
			}else{
				track = append(track, nums[i])
				used[i]=true
				//进入下一组决策
				backtrack(nums,track)
				track=track[:len(track)-1]
				used[i]=false
			}
		}
	}
	backtrack(nums,track)
	return res
}
func main() {
	input:=[]int{1,2,3}
	ret:=permute(input)
	fmt.Println(ret)
}
