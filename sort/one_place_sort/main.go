//题目描述
//给定一个非空数组（列表），其元素数据类型为整型，请按照数组元素十进制最低位从小到大进行排序，十进制最低位相同的元素，相对位置保持不变。
//当数组元素为负值时，十进制最低位等同于去除符号位后对应十进制值最低位。
//输入描述
//给定一个非空数组，其元素数据类型为32位有符号整数，数组长度[1, 1000]
//输出描述
//输出排序后的数组
//用例
//输入	1,2,5,-21,22,11,55,-101,42,8,7,32
//输出	1,-21,11,-101,2,22,42,32,5,55,7,8
//说明	无

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type SortedWithIndex struct {
	Num       int
	LastDigit int
	Index     int
}

func sortWithOnePlace(arr []int) string {
	if len(arr) <= 0 {
		return ""
	}

	indexedArr := make([]SortedWithIndex, len(arr))
	for i, num := range arr {
		absNum := num
		if num < 0 {
			absNum = -num
		}

		lastDigit := absNum % 10
		indexedArr[i] = SortedWithIndex{
			Num:       num,
			LastDigit: lastDigit,
			Index:     i,
		}
	}

	// 进行排序
	sort.SliceStable(indexedArr, func(i, j int) bool {
		if indexedArr[i].LastDigit != indexedArr[j].LastDigit {
			return indexedArr[i].LastDigit < indexedArr[j].LastDigit
		}

		return indexedArr[i].Index < indexedArr[j].Index
	})

	// 获取排序后的数组
	sorted := make([]string, len(indexedArr))
	for i := 0; i < len(indexedArr); i++ {
		sorted[i] = strconv.Itoa(indexedArr[i].Num)
	}

	return strings.Join(sorted, ",")
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
		res := sortWithOnePlace(arr)
		fmt.Println(res)
	}
}
