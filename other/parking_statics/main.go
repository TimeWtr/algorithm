package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handler(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	total, i := 0, 0
	for i < len(arr) {
		// 跳过空闲车位
		if arr[i] == 0 {
			i++
			continue
		}

		// 统计连续停车的车位长度
		length := 0
		for i < len(arr) && arr[i] == 1 {
			length++
			i++
		}
		total += minStatic(length)
	}

	return total
}

func minStatic(length int) int {
	trucks := length / 3
	remainders := length % 3
	if remainders == 0 {
		return trucks
	} else if remainders == 1 {
		// 停小车
		return trucks + 1
	} else {
		// 停货车
		return trucks + 1
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arrStr := strings.Split(strings.TrimSpace(scanner.Text()), ",")
		arr := make([]int, 0, len(arrStr))
		for _, str := range arrStr {
			val, _ := strconv.Atoi(str)
			arr = append(arr, val)
		}

		cnt := handler(arr)
		fmt.Println(cnt)
	}
}
