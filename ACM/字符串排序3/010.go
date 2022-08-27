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
		strs := strings.Split(input.Text(), ",")
		sort.Strings(strs)
		for i := 0; i < len(strs)-1; i++ {
			fmt.Print(strs[i], ",")
		}
		fmt.Println(strs[len(strs)-1])
	}
}
