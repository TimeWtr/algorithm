package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func handler(str string) string {
	if len(str) == 0 {
		return ""
	}

	if isMiniStr(str) {
		return str
	}

	// 构建最小序
	miniStr := buildMiniStr(str)
	// 找出第一个不一致的字符下标
	needSwapIndex := -1
	for i := 0; i <= len(miniStr)-1; i++ {
		if miniStr[i] != str[i] {
			needSwapIndex = i
			break
		}
	}

	// 从原始字符串中找需要交换的目标字符的下标
	targetStr := miniStr[needSwapIndex]
	swapIndex := -1
	for i := len(str) - 1; i >= 0; i-- {
		if targetStr == str[i] {
			swapIndex = i
			break
		}
	}

	// 交换
	chars := []byte(str)
	chars[needSwapIndex], chars[swapIndex] = chars[swapIndex], chars[needSwapIndex]

	return string(chars)
}

func buildMiniStr(str string) string {
	chars := strings.Split(str, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

// 判断是否是最小序字典
func isMiniStr(str string) bool {
	for i := 1; i < len(str); i++ {
		if str[i] < str[i-1] {
			return false
		}
	}

	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := strings.TrimSpace(scanner.Text())
		res := handler(str)
		fmt.Println(res)
	}
}
