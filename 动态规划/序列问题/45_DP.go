package leetcode

// func jump(nums []int) int {
// 	// 状态：f[i] 表示从起点到当前位置最小次数
// 	// 推导：f[i] = f[j],a[j]+j >=i,min(f[j]+1)
// 	// 初始化：f[0] = 0
// 	// 结果：f[n-1]
// 	dp := make([]int, len(nums))

// 	dp[0] = 0
// 	for i := 1; i < len(nums); i++ {
// 		//f[i]的最大值为i
// 		dp[i] = i
// 		//遍历之前的结果去一个最小值+1
// 		for j := 0; j < i; j++ {
// 			if nums[j]+j >= i {
// 				dp[i] = min(dp[j]+1,dp[i])
// 			}
// 		}
// 	}
// 	return dp[len(nums)-1]
// }
// func min(a, b int) int{
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

func jump(nums []int) int {
	n:=len(nums)
	// 备忘录都初始化为 n，相当于 INT_MAX
    // 因为从 0 跳到 n - 1 最多 n - 1 步
	memo:=make([]int,n)
	for i, _ := range memo {
		memo[i]=n
	}

	return dp(nums,0)
}
func dp(nums []int)  {
	
}
