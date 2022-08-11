package leetcode

func uniquePaths_1(m int, n int) int {
	return help_1(m-1,n-1)
}
//从0,0走到x,y 有多少路径
func help_1(x,y int)int{
	if x==0||y==0{
		return 1
	}
	return help_1(x-1,y)+help_1(x,y-1)
}