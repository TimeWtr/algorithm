package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(subStr string, str string) int {
	i, j := 0, 0
	// 终止条件：子串/完整字符串任意一个遍历完成
	for i < len(subStr) && j < len(str) {
		if subStr[i] == str[j] {
			i++
		}
		j++
	}

	if i == len(subStr) {
		// 找到了
		return j - 1
	}

	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		subStr := scanner.Text()

		if !scanner.Scan() {
			break
		}
		str := scanner.Text()
		pos := check(subStr, str)
		fmt.Println(pos)
	}
}
