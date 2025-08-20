package main

import (
	"bufio"
	"fmt"
	"os"
)

func bracketDepths(str string) int {
	if len(str)%2 != 0 {
		return 0
	}

	maxDepth := 0
	var stack []rune
	for _, r := range str {
		// 左侧全部入栈
		if r == '(' || r == '[' || r == '{' {
			stack = append(stack, r)
			maxDepth = max(maxDepth, len(stack))
			continue
		}

		n := len(stack)
		if r == ')' {
			if n > 0 && stack[n-1] == '(' {
				stack = stack[:n-1]
			} else {
				return 0
			}
		}

		if r == ']' {
			if n > 0 && stack[n-1] == '[' {
				stack = stack[:n-1]
			} else {
				return 0
			}
		}

		if r == '}' {
			if n > 0 && stack[n-1] == '{' {
				stack = stack[:n-1]
			} else {
				return 0
			}
		}
	}

	if len(stack) == 0 {
		return maxDepth
	}

	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		depths := bracketDepths(str)
		fmt.Println("max depths: ", depths)
	}
}
