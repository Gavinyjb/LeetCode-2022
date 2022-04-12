package main

import "fmt"

func isPerfectSquare(num int) bool {
	if num==0{
		return false
	}
	if num==1{
		return true
	}
	low,high:=1,num
	for low<high{
		mid:=low+(high-low)>>1
		if mid*mid==num{
			return true
		}else if mid*mid>num {
			high=mid
		}else {
			low=mid+1
		}
	}
	return false
}

func main(){
	ret:=isPerfectSquare(1)
	fmt.Println(ret)
}
