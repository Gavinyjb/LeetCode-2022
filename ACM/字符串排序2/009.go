package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := strings.Split(input.Text(), " ")
		sort.Strings(nums)
		for _, num := range nums {
			fmt.Print(num, " ")
		}
		fmt.Println()
	}
}
