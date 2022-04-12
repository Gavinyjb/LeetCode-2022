package main
func mySqrt(x int) int {
	// 解法一 二分, 找到最后一个满足 n^2 <= x 的整数n
	l,r:=0,x
	for l<r{
		mid:=l+(r-l)>>1
		if mid*mid>x{
			r=mid-1
		}else {
			l=mid
		}
	}
	return l
}
func main() {

}
