package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func KLettersLengths(str string, k int) int {
	// 边界条件
	if k <= 0 {
		return -1
	}

	cm := make(map[byte]int)
	if len(str) > 0 {
		prev := str[0]
		cnt := 1

		for i := 1; i < len(str); i++ {
			if str[i] == prev {
				cnt++
			} else {
				if count, ok := cm[prev]; !ok || cnt > count {
					cm[prev] = cnt
				}
				prev = str[i]
				cnt = 1
			}
		}

		// 处理最后一段连续字母的问题
		if count, ok := cm[prev]; !ok || cnt > count {
			cm[prev] = cnt
		}
	}

	// 边界条件
	if len(str) < k {
		return -1
	}
	// 转换成数组
	arr := make([]int, 0, len(cm))
	for _, v := range cm {
		arr = append(arr, v)
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	return arr[k-1]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		str := scanner.Text()

		if !scanner.Scan() {
			break
		}
		k, _ := strconv.Atoi(scanner.Text())
		length := KLettersLengths(str, k)
		fmt.Println(length)
	}
}
