package main

import "fmt"

//4
//abcd
//efgh

//aebfcgdh

func process(str1, str2 string) string {
	if len(str1) != len(str2) {
		return ""
	}
	s1 := []byte(str1)
	s2 := []byte(str2)
	ret := make([]byte, 0)

	for i := 0; i < len(s1); i++ {
		ret = append(ret, s1[i], s2[i])
	}
	return string(ret)
}
func main() {
	n := 0
	fmt.Scanf("%d\n", &n)
	str1, str2 := "", ""
	fmt.Scan(&str1)
	fmt.Scan(&str2)

	res := process(str1, str2)
	fmt.Println(res)
}
