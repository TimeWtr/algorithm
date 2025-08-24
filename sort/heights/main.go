//题目描述
//小明今年升学到了小学1年级来到新班级后，发现其他小朋友身高参差不齐，然后就想基于各小朋友和自己的身高差，对他们进行排序，请帮他实现排序。
//
//输入描述
//第一行为正整数 h和n，0<h<200 为小明的身高，0<n<50 为新班级其他小朋友个数。
//
//第二行为n个正整数，h1 ~ hn分别是其他小朋友的身高，取值范围0<hi<200，且n个正整数各不相同。
//
//输出描述
//输出排序结果，各正整数以空格分割，
//
//和小明身高差绝对值最小的小朋友排在前面，
//
//和小明身高差绝对值最大的小朋友排在后面，
//
//如果两个小朋友和小明身高差一样，则个子较小的小朋友排在前面。
//
//用例
//输入	100 10
//95 96 97 98 99 101 102 103 104 105
//输出	99 101 98 102 97 103 96 104 95 105
//说明	小明身高100，班级学生10个，身高分别为95 96 97 98 99 101 102 103 104 105，按身高差排序后结果为：99 101 98 102 97 103 96 104 95 105。

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sorted(height int, arr []int) string {
	if len(arr) == 0 {
		return ""
	}

	// 进行排序
	sort.Slice(arr, func(i, j int) bool {
		// 计算绝对差值
		// 主逻辑：比较绝对差值
		deltaI := abs(arr[i], height)
		deltaJ := abs(arr[j], height)
		if deltaI != deltaJ {
			return deltaI < deltaJ
		}

		// 绝对差值相等的情况下比较身高的大小
		return arr[i] < arr[j]
	})

	strArr := make([]string, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		strArr = append(strArr, strconv.Itoa(arr[i]))
	}

	return strings.Join(strArr, " ")
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		firstStrArr := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		height, _ := strconv.Atoi(firstStrArr[0])
		number, _ := strconv.Atoi(firstStrArr[1])

		if !scanner.Scan() {
			break
		}
		secondArr := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		if len(secondArr) < number {
			fmt.Println("")
			break
		}
		heights := make([]int, 0, len(secondArr))
		for _, val := range secondArr {
			h, _ := strconv.Atoi(val)
			heights = append(heights, h)
		}
		res := sorted(height, heights)
		fmt.Println(res)
	}
}
