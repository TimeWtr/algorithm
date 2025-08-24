package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func calculate(x, y int) int {
	// 边界条件
	if x <= 0 || y <= 0 || y >= 5 {
		return -1
	}

	// 计算字母的组合数
	lettersCombinations := math.Pow(26, float64(y))
	// 计算需要的组合数
	neededCombinations := float64(x) / lettersCombinations
	if neededCombinations <= 1 {
		return 1
	}

	z := math.Ceil(math.Log10(neededCombinations))
	if z < 1 {
		z = 1
	}

	return int(z)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		x, err := strconv.Atoi(arr[0])
		y, err := strconv.Atoi(arr[1])
		if err != nil {
			fmt.Println(-1)
			return
		}
		res := calculate(x, y)
		fmt.Println(res)
	}
}
