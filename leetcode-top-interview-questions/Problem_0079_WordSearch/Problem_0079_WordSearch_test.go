package Problem_0079_WordSearch

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	//board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	board := [][]byte{{'C', 'A', 'A'}, {'A', 'A', 'A'}, {'B', 'C', 'D'}}
	//board := [][]byte{{'A', 'A', 'A', 'A', 'A', 'A'}, {'A', 'A', 'A', 'A', 'A', 'A'}, {'A', 'A', 'A', 'A', 'A', 'A'}, {'A', 'A', 'A', 'A', 'A', 'A'}, {'A', 'A', 'A', 'A', 'A', 'B'}, {'A', 'A', 'A', 'A', 'B', 'A'}}
	//word := "ABCCED"
	word := "AAB"
	//word := "AAAAAAAAAAAAABB"
	fmt.Println(exist(board, word))
}
func TestTemp(t *testing.T) {
	choose := [][]int{
		{0, 1},  //右
		{0, -1}, //左
		{1, 0},  //下
		{-1, 0}, //上
	}
	for i := 0; i < len(choose); i++ {
		fmt.Print(choose[i][0], choose[i][1], " ")
	}
	fmt.Println("-------------")
	for _, ints := range choose {
		fmt.Print(ints[0], ints[1], " ")
	}
	fmt.Println("-------------")
}
