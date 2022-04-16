package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// NestedInteger define
type NestedInteger struct {
	Num  int
	List []*NestedInteger
}

// IsInteger define
func (n NestedInteger) IsInteger() bool {
	if n.List == nil {
		return true
	}
	return false
}

// GetInteger define
func (n NestedInteger) GetInteger() int {
	return n.Num
}

// SetInteger define
func (n *NestedInteger) SetInteger(value int) {
	n.Num = value
}

// Add define
func (n *NestedInteger) Add(elem NestedInteger) {
	n.List = append(n.List, &elem)
}

// GetList define
func (n NestedInteger) GetList() []*NestedInteger {
	return n.List
}

// Print define
func (n NestedInteger) Print() {
	if len(n.List) != 0 {
		for _, v := range n.List {
			if len(v.List) != 0 {
				v.Print()
				return
			}
			fmt.Printf("%v ", v.Num)
		}
	} else {
		fmt.Printf("%v ", n.Num)
	}
	fmt.Printf("\n")
}

func deserialize(s string) *NestedInteger {
	stack, cur := []*NestedInteger{}, &NestedInteger{}
	for i := 0; i < len(s); {
		switch {
		case s[i]) || s[i] == '-':
			j := 0
			for j = i + 1; j < len(s) && isDigital(s[j]); j++ {
			}
			num, _ := strconv.Atoi(s[i:j])
			next := &NestedInteger{}
			next.SetInteger(num)
			if len(stack) > 0 {
				stack[len(stack)-1].List = append(stack[len(stack)-1].GetList(), next)
			} else {
				cur = next
			}
			i = j
		case s[i] == '[':
			next := &NestedInteger{}
			if len(stack) > 0 {
				stack[len(stack)-1].List = append(stack[len(stack)-1].GetList(), next)
			}
			stack = append(stack, next)
			i++
		case s[i] == ']':
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			i++
		case s[i] == ',':
			i++
		}
	}
	return cur
}
func main() {

}
