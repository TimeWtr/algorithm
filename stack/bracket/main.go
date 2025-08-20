package main

import (
	"bufio"
	"fmt"
	"os"
)

func bracketMatch(str string) bool {
	var stack []rune
	for _, r := range str {
		if r == '(' || r == '[' || r == '{' {
			stack = append(stack, r)
			continue
		}

		n := len(stack)
		if r == ')' {
			if n > 0 && stack[n-1] == '(' {
				stack = stack[:n-1]
			} else {
				return false
			}
		}

		if r == ']' {
			if n > 0 && stack[n-1] == '[' {
				stack = stack[:n-1]
			} else {
				return false
			}
		}

		if r == '}' {
			if n > 0 && stack[n-1] == '{' {
				stack = stack[:n-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		matched := bracketMatch(str)
		fmt.Println("是否成对匹配：", matched)
	}
}
