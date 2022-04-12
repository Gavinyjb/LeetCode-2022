package main

import (
	"fmt"
	"sort"
)

func permute(nums []int) [][]int {
	sort.Ints(nums)
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
			}else if i > 0 && nums[i] == nums[i - 1] && !used[i - 1] {
				continue
			} else{
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
	input:=[]int{1,1,2}
	ret:=permute(input)
	fmt.Println(ret)

}
