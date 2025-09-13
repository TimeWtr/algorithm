// 题目：荒岛求生
//题目描述
//一个荒岛上有若干人，岛上只有一条路通往岛屿两端的港口，大家需要逃往两端的港口才可逃生。
//
//假定每个人移动的速度一样，且只可选择向左或向右逃生。
//
//若两个人相遇，则进行决斗，战斗力强的能够活下来，并损失掉与对方相同的战斗力；若战斗力相同，则两人同归于尽。
//
//输入描述
//给定一行非 0 整数数组，元素个数不超过30000；
//
//正负表示逃生方向（正表示向右逃生，负表示向左逃生），绝对值表示战斗力，越左边的数字表示里左边港口越近，逃生方向相同的人永远不会发生决斗。
//
//输出描述
//能够逃生的人总数，没有人逃生输出0，输入异常时输出-1。
//
//示例1
//输入
//
//5 10 8 -8 -5
//1
//输出
//
//2
//1
//说明
//
//第3个人和第4个人同归于尽，第2个人杀死第5个人并剩余5战斗力，第1个人没有遇到敌人。

// 题解：使用栈来进行对比，如果是正数直接入栈，负数则直接与栈顶元素对比

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getIntArr(strArr []string) []int {
	if len(strArr) == 0 {
		return nil
	}

	arr := make([]int, len(strArr))
	for i := range strArr {
		arr[i], _ = strconv.Atoi(strArr[i])
	}

	return arr
}

func Survives(arr []int) int {
	var stack []int
	for i := range arr {
		curr := arr[i]
		if curr > 0 {
			stack = append(stack, curr)
			continue
		}

		survive := true
		for len(stack) > 0 && stack[len(stack)-1] > 0 {
			// 栈内有人且栈顶的人武力值不为负（向右）
			top := stack[len(stack)-1]
			if top == -curr {
				// 同归于尽了
				survive = false
				stack = stack[:len(stack)-1]
				break
			} else if top > -curr {
				// 向右的人武力值更高
				stack[len(stack)-1] = top + curr
				survive = false
				break
			} else {
				// 向左的武力值更高
				curr = top + curr
				stack = stack[:len(stack)-1]
				if curr == 0 {
					survive = false
					break
				}
			}
		}

		if survive {
			stack = append(stack, curr)
		}
	}

	return len(stack)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.Fields(strings.TrimSpace(scanner.Text()))
		arr := getIntArr(line)
		count := Survives(arr)
		fmt.Println(count)
	}
}
