package main
func moveZeroes(nums []int)  {
	if len(nums)==0{
		return
	}
	slow,fast:=0,0
	for ;fast<len(nums);fast++{
		if nums[fast]!=0{
			nums[slow],nums[fast]=nums[fast],nums[slow]
			slow++
		}
	}
	return
}
func main() {

}
