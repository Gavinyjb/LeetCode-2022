package leetcode


func climbStairs(n int) int {
	return help_1(n)
}
//剩余n 级台阶
func help_1(n int)int{
	if n==1||n==0{
		return 1
	}
	return help_1(n-1)+help_1(n-2)
}