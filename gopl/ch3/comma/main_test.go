package main

import (
	"fmt"
	"testing"
)

func TestComma(t *testing.T) {
	strList := []string{"1", "12", "123", "1234", "1234567890"}
	for _, s := range strList {
		fmt.Println(comma(s))
		fmt.Println(comma1(s))
		if comma1(s) != comma(s) {
			t.Errorf("递归改迭代出错！")
		}
	}
}
func BenchmarkComma(b *testing.B) {
	strList := []string{"1", "12", "123", "1234", "1234567890"}
	for i := 0; i < b.N; i++ {
		for _, s := range strList {
			comma(s)
		}
	}
}
