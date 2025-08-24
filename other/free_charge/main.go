package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handler(group map[string][]int) int {
	if len(group) == 0 {
		return 0
	}

	totalCount := 0
	for _, millsArr := range group {
		miniMill := millsArr[0]
		// 找出当前分组中最小的毫秒
		for _, mill := range millsArr {
			if mill < miniMill {
				miniMill = mill
			}
		}

		// 计算当前分组中处于最小毫秒的计数
		for _, mill := range millsArr {
			if mill == miniMill {
				totalCount++
			}
		}

	}

	return totalCount
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		n, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil || n > 50000 || n <= 0 {
			fmt.Println(-1)
			return
		}

		group := map[string][]int{}
		for i := 0; i < n; i++ {
			if !scanner.Scan() {
				break
			}
			str := strings.TrimSpace(scanner.Text())
			if len(str) != 23 {
				fmt.Println(-1)
				break
			}
			groupID := str[:19]
			mills, _ := strconv.Atoi(str[20:])
			if mills > 999 || mills < 0 {
				fmt.Println(-1)
				break
			}
			arr, ok := group[groupID]
			if !ok {
				var newArr []int
				newArr = append(newArr, mills)
				group[groupID] = newArr
			} else {
				arr = append(arr, mills)
				group[groupID] = arr
			}
		}

		cnt := handler(group)
		fmt.Println(cnt)
	}
}
