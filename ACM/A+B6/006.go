package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//https://ac.nowcoder.com/acm/contest/5652/F
func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := strings.Split(input.Text(), " ")
		res := 0
		for i := 1; i < len(nums); i++ {
			num, _ := strconv.Atoi(nums[i])
			res += num
		}
		fmt.Println(res)
	}
}
