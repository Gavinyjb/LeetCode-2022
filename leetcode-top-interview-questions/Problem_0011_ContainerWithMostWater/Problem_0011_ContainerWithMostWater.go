package leetcodetopinterviewquestions

func maxArea(height []int) int {
	res:=0
	l,r:=0,len(height)-1
	for l<r{
		res=max(res,(r-l)*min(height[l],height[r]))
		if height[l]<height[r]{
			l++
		}else {
			r--
		}
	}
	return res
}
func max(a , b int) int  {
	if a<b{
		return b
	}
	return a
}
func min(a,b int)int  {
	if a<b{
		return a
	}
	return b
}