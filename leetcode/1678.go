package leetcode

import (
	"strings"
)

func interpret(command string) string {
	prev := 0
	res := strings.Builder{}
	for i := range command {
		switch command[prev : i+1] {
		case "G":
			res.WriteRune('G')
		case "()":
			res.WriteRune('o')
		case "(al)":
			res.WriteString("al")
		}
		if command[prev:i+1] == "G" || command[prev:i+1] == "()" || command[prev:i+1] == "(al)" {
			prev = i + 1
		}
	}
	return res.String()
}
