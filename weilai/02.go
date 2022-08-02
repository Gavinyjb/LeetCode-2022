package main

import "fmt"

func main() {
	length := 0
	fmt.Scanf("%d\n", &length)
	str := ""
	fmt.Scanf("%v\n", &str)
	//fmt.Println(str)
	res := make([]byte, 0)
	cnt := 0
	flag := false
	for i := 0; i < len(str); i++ {
		if len(res) > 0 {
			if res[len(res)-1] == '(' && str[i] == ')' {
				res = res[:len(res)-1]
				cnt += 2
			} else {
				res = append(res, str[i])
			}
		} else {
			if i != 0 {
				flag = true
			}
			res = append(res, str[i])
		}
	}
	if flag {
		fmt.Println(cnt)
	} else {
		fmt.Println(cnt)
	}
}
