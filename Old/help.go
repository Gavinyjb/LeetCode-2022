package main

import (
	"fmt"
	"sort"
)
func help(nums []int,l,r int,sum)int  {
	ret:=0
	for l<r{
		if nums[l]+nums[r]<sum{
			l++
		}
		if nums[l]+nums[r]>sum{
			r--
		}
		if nums[l]+nums[r]==sum{
			ret++
		}
	}
	for l+1<r{
		if nums[l]==
	}
}


func main() {
	n := 0
	fmt.Scan(&n)
	nums:=make([]int,n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i]<nums[j]
	})
	// i+k=3j
	i,k:=0,len(nums)-1


	fmt.Println(result)
}