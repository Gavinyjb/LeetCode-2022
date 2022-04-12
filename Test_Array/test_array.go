package main

import "fmt"

func testArray()  {
	var array [2][3]int
	array[0]=[3]int{1,2,3}
	array[1]=[3]int{4,5,6}
	fmt.Println(&array[0][0], &array[0][1] ,&array[0][2] )
	fmt.Println(&array[1][0], &array[1][1],&array[1][2])
}

func main() {
	testArray()
}
