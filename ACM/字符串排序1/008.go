package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 0
	fmt.Scan(&n)
	strList := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&strList[i])
	}
	sort.Strings(strList)
	for i := 0; i < n; i++ {
		fmt.Print(strList[i], " ")
	}
}
