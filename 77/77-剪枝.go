package main

import "fmt"

func combine(n int, k int) [][]int {
	nums:=make([]int,n)
	for i := 0; i < n; i++ {
		nums[i]=i+1
	}
	res:=make([][]int,0)
	track:=make([]int,0)
	var backtrack func(track []int,start int)
	backtrack= func(track []int, start int) {
		//终止条件
		if len(track)==k{
			temp:=make([]int,k)
			copy(temp,track)
			res = append(res, temp)
		}
		if len(track)+n-start+1<k{
			return
		}
		//for 选择列表 组合问题，选过了的不再重复选
		for i := start; i <n; i++ {
			//做选择
			track = append(track, nums[i])
			//进入下一层决策树
			backtrack(track,i+1)
			track=track[:len(track)-1]
		}
	}
	backtrack(track,0)
	return res
}
func main() {
	fmt.Println(combine(4,2))
}
