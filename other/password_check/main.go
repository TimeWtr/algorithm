package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(originPassword string) string {
	if len(originPassword) == 0 {
		return ""
	}

	var stack []byte
	for i := 0; i < len(originPassword)-1; i++ {
		if originPassword[i] == ' ' {
			continue
		}

		if originPassword[i] == '<' {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, originPassword[i])
	}

	if len(stack) == 0 {
		return ""
	}

	var password string
	lower, upper := 0, 0
	number, other := 0, 0
	for _, char := range stack {
		if char >= 'A' && char <= 'Z' {
			upper++
		} else if char >= 'a' && char <= 'z' {
			lower++
		} else if char >= '0' && char <= '9' {
			number++
		} else {
			other++
		}
		password += string(char)
	}

	valid := lower >= 1 && upper >= 1 && number >= 1 && other >= 1
	return fmt.Sprintf("%s,%v", password, valid)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		originalPassword := scanner.Text()
		res := check(originalPassword)
		fmt.Println(res)
	}
}
