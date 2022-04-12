package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {

	sort.Ints(nums)  // 这里是去重的关键逻辑
	res:=make([][]int,0)
	track:=make([]int,0)

	var backtrack func(nums []int,track []int,start int)
	backtrack= func(nums []int, track []int,start int) {
		temp:=make([]int,len(track))
		copy(temp,track)
		res = append(res, temp)

		//for 选择 in 选择列表:
		for i := start; i < len(nums); i++ {
			//做选择
			if (i>start&&nums[i]==nums[i-1]){
				continue
			}
			track = append(track, nums[i])
			backtrack(nums,track,i+1)
			track=track[:len(track)-1]
		}
	}
	backtrack(nums,track,0)

	return res

}


func main() {
	input:=[]int{1,2,2}
	ret:=subsetsWithDup(input)
	fmt.Println(ret)
}
