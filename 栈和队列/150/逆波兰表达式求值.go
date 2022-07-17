package leetcode

import (
	"log"
	"strconv"
)

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return -1
	}
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+", "-", "*", "/":
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var result int
			switch tokens[i] {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				result = a / b
			}
			stack = append(stack, result)
		default:
			val, err := strconv.Atoi(tokens[i])
			if err != nil {
				log.Println(err)
			}
			stack = append(stack, val)
		}
	}
	return stack[0]
}
