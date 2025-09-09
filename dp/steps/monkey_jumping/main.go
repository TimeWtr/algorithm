// 题目：猴子跳台阶
// 题目描述
// 一天一只顽猴想去从山脚爬到山顶，途中经过一个有个N个台阶的阶梯，但是这猴子有一个习惯：

// 每一次只能跳1步或跳3步，试问猴子通过这个阶梯有多少种不同的跳跃方式？

// 输入描述
// 输入只有一个整数N（0<N<=50）此阶梯有多少个台阶。

// 输出描述
// 输出有多少种跳跃方式（解决方案数）。

// 用例1
// 输入
// 50
// 输出
// 122106097
// 用例2
// 输入
// 3
// 输出
// 2

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handler(steps int) int {
	if steps < 0 {
		return -1
	}

	dp := make([]int, steps+1)
	dp[0] = 1
	if steps >= 1 {
		dp[1] = 1
	}

	if steps >= 2 {
		dp[2] = 1
	}

	for i := 3; i <= steps; i++ {
		dp[i] = dp[i-1] + dp[i-3]
	}

	return dp[steps]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		steps, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		res := handler(steps)
		fmt.Println(res)
	}
}
