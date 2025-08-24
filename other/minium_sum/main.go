// 题目描述
// 给定一个从小到大的有序整数序列（存在正整数和负整数）数组 nums ，请你在该数组中找出两个数，其和的绝对值(|nums[x]+nums[y]|)为最小值，并返回这个绝对值。
// 每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
//
// 输入描述
// 一个通过空格分割的有序整数序列字符串，最多1000个整数，且整数数值范围是 -65535~65535。
//
// 输出描述
// 两数之和绝对值最小值
//
// 用例
// 输入	-3 -1 5 7 11 15
// 输出	2
// 说明	因为 |nums[0] + nums[2]| = |-3 + 5| = 2 最小，所以返回 2。
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	left, right := 0, len(arr)-1
	minSum := abs(arr[left] + arr[right])
	for left < right {
		sum := arr[left] + arr[right]
		absSum := abs(sum)
		if absSum < minSum {
			minSum = absSum
		}

		if sum > 0 {
			right--
		} else if sum < 0 {
			left++
		} else {
			return 0
		}
	}

	return minSum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strArr := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		arr := make([]int, len(strArr))
		for i, str := range strArr {
			arr[i], _ = strconv.Atoi(str)
		}
		minSum := calculate(arr)
		fmt.Println(minSum)
	}
}
