package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func findLetters(subStr, str string) string {
	if len(subStr) == 0 || len(str) == 0 || len(subStr) > len(str) {
		return ""
	}

	// 定义一个集合
	existMap := make(map[byte]struct{})
	mp := make(map[byte]struct{}, len(subStr))
	for _, v := range subStr {
		mp[byte(v)] = struct{}{}
	}

	for _, v := range str {
		if _, ok := mp[byte(v)]; ok {
			existMap[byte(v)] = struct{}{}
		}
	}

	if len(existMap) == 0 {
		return ""
	}

	arr := make([]byte, 0, len(existMap))
	for k, _ := range existMap {
		arr = append(arr, k)
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	return string(arr)
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
		res := findLetters(subStr, str)
		fmt.Println(res)
	}
}
