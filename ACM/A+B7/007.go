package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//https://ac.nowcoder.com/acm/contest/5652/G
func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := strings.Split(input.Text(), " ")
		sum := 0
		for i, _ := range nums {
			num, _ := strconv.Atoi(nums[i])
			sum += num
		}
		fmt.Println(sum)
	}
}
