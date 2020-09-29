package leetcode

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	l := []int{}
	res := 0
	for _, val := range tokens {
		a, b := 0, 0
		res = 0
		if isOperand(val) {
			a = l[len(l)-2]
			b = l[len(l)-1]
		}
		switch val {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			res = a / b
		default:
			num, _ := strconv.Atoi(val)
			l = append(l, num)
		}

		if isOperand(val) {
			l = l[:len(l)-2]
			l = append(l, res)
		}
	}
	if len(tokens) == 1 {
		num, _ := strconv.Atoi(tokens[0])
		return num
	}
	return res
}

func isOperand(c string) bool {
	return c == "+" || c == "-" || c == "*" || c == "/"
}
