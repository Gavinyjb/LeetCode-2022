package main
//有A:x B:y 最多有几个礼盒
func help(x,y int) int {
	//无法制作
	if x<=0||y<=0||x+y<=2{
		return 0
	}
	//只能做一个
	if x==1&&y>=2{
		return 1
	}
	if y==1&&x>=2{
		return 1
	}
	return 1+max(help(x-1,y-2),help(x-2,y-1))
}
func max(a,b int)int  {
	if a<b{
		return b
	}
	return a
}
func main(){
	
}