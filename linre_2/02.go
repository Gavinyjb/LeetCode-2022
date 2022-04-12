package main

import (
	"fmt"
	"sort"
)

type peoples struct {
	height int
	name  string
}
func main() {
	var n int  //人数
	fmt.Scanf("%d", &n)
	p:=make([]peoples,n)

	for i := 0; i < n; i++ {
		fmt.Scan(&p[i].height)
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&p[i].name)
	}
	sort.Slice(p, func(i, j int) bool {
		if p[i].height<p[j].height{
			return true
		}else if p[i].height==p[j].height{
			if p[i].name<p[j].name{
				return true
			}
		}
		return false
	})
	for _, v := range p {
		fmt.Printf("%v ",v.name)
	}
	fmt.Println()

}