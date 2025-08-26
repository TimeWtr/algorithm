package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var res [][]int

func recr(n, status int, arr []int) []int {
	if n <= 1 {
		return arr
	}

	for i := 1; i <= int(math.Floor(float64(n/2))); i++ {
		if i%2 == 0 {
			// 偶数
			if status == 1 {
				continue
			}

			if status == 0 && arr[len(arr)-1]%2 != 0 {
				arr = append(arr, i)
				res = append(res, arr)
				recr(i, status, arr)
			}

			if status == 2 {
				arr = append(arr, i)
				res = append(res, arr)
				recr(i, status, arr)
			}
		} else {
			if status == 2 {
				continue
			}

			if status == 0 && arr[len(arr)-1]%2 == 0 {
				recr(i, status, arr)
			}

			if status == 1 {
				recr(i, status, arr)
			}
		}
	}
	return arr
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if n < 1 || n > 500 {
			fmt.Println(-1)
			break
		}

		arr := []int{n}
		for i := 0; i < 3; i++ {
			res = append(res, recr(n, i, arr))
		}
		fmt.Println(res)
	}
}
