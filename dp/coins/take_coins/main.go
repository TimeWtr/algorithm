// 题目：拿硬币
// 桌上有 n 堆力扣币，每堆的数量保存在数组 coins 中。我们每次可以选择任意一堆，拿走其中的一枚或者两枚，
// 求拿完所有力扣币的最少次数。

// 示例 1：

// 输入：[4,2,1]

// 输出：4

// 解释：第一堆力扣币最少需要拿 2 次，第二堆最少需要拿 1 次，第三堆最少需要拿 1 次，总共 4 次即可拿完。

// 示例 2：

// 输入：[2,3,10]

// 输出：8

// 限制：

// 1 <= n <= 4
// 1 <= coins[i] <= 10

// 解题思路：
// 1. 4枚：每次拿2枚，最少拿2次
// 2. 3枚：第一次拿2枚，第二次拿1枚，最少拿2次
// 3. 2枚：一次拿2枚，最少拿1次
// 4. 1枚：一次拿1枚，最少1次
// 总结：最少拿次数是c/2，c是硬币的数量，利用向上取整，计算公式(c+1)/2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getStringArr(str string) []string {
	if str == "" {
		return nil
	}

	oriStr := strings.TrimRight(strings.TrimLeft(strings.TrimSpace(str), "["), "]")
	return strings.Split(oriStr, ",")
}

func getIntArr(strArr []string) []int {
	if len(strArr) == 0 {
		return nil
	}

	arr := make([]int, 0, len(strArr))
	for _, v := range strArr {
		val, _ := strconv.Atoi(v)
		arr = append(arr, val)
	}

	return arr
}

func taker(arr []int) int {
	sum := 0
	for _, v := range arr {
		minCount := (v + 1) / 2
		sum += minCount
	}

	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		strArr := getStringArr(str)
		arr := getIntArr(strArr)
		count := taker(arr)
		fmt.Println(count)
	}
}
