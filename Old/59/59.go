package main

import "fmt"

func generateMatrix(n int) [][]int {
	if n==0{
		return nil
	}
	res:= make([][]int, n)
	for i := 0; i < n; i++ {
		res[i]=make([]int,n)
	}
	//fmt.Println(res)
	num := 1
	top, bottom := 0, n-1
	left, right := 0, n-1
	for num <= n*n {
		//顶部
		for i := left; i <= right; i++ {
			res[top][i] = num
			num++
		}
		top++
		//右边
		for i := top; i <=bottom; i++ {
			res[i][right]=num
			num++
		}
		right--
		//底部
		for i := right; i >=left ; i-- {
			res[bottom][i]=num
			num++
		}
		bottom--
		//左侧
		for i := bottom; i >=top ; i-- {
			res[i][left]=num
			num++
		}
		left++
	}
	return res
}
//func generateMatrix(n int) [][]int {
//	top, bottom := 0, n-1
//	left, right := 0, n-1
//	num := 1
//	tar := n * n
//	matrix := make([][]int, n)
//	for i := 0; i < n; i++ {
//		matrix[i] = make([]int, n)
//	}
//	for num <= tar {
//		for i := left; i <= right; i++ {
//			matrix[top][i] = num
//			num++
//		}
//		top++
//		for i := top; i <= bottom; i++ {
//			matrix[i][right] = num
//			num++
//		}
//		right--
//		for i := right; i >= left; i-- {
//			matrix[bottom][i] = num
//			num++
//		}
//		bottom--
//		for i := bottom; i >= top; i-- {
//			matrix[i][left] = num
//			num++
//		}
//		left++
//	}
//	return matrix
//}
func main() {
	ret:=generateMatrix(5)
	fmt.Println(ret)
}
