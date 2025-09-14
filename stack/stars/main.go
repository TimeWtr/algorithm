// 题目：小行星碰撞
// 给定一个整数数组 asteroids，表示在同一行的小行星。数组中小行星的索引表示它们在空间中的相对位置。
//
//对于数组中的每一个元素，其绝对值表示小行星的大小，正负表示小行星的移动方向（正表示向右移动，负表示向左移动）。每一颗小行星以相同的速度移动。
//
//找出碰撞后剩下的所有小行星。碰撞规则：两个小行星相互碰撞，较小的小行星会爆炸。如果两颗小行星大小相同，则两颗小行星都会爆炸。两颗移动方向相同的小行星，永远不会发生碰撞。
//
//
//
//示例 1：
//
//输入：asteroids = [5,10,-5]
//输出：[5,10]
//解释：10 和 -5 碰撞后只剩下 10 。 5 和 10 永远不会发生碰撞。
//示例 2：
//
//输入：asteroids = [8,-8]
//输出：[]
//解释：8 和 -8 碰撞后，两者都发生爆炸。
//示例 3：
//
//输入：asteroids = [10,2,-5]
//输出：[10]
//解释：2 和 -5 发生碰撞后剩下 -5 。10 和 -5 发生碰撞后剩下 10 。

// 题解：使用栈来存储向右的方向的行星，遇到向左的行星则直接碰撞

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

	intArr := make([]int, len(strArr))
	for i := range strArr {
		intArr[i], _ = strconv.Atoi(strArr[i])
	}
	return intArr
}

func Survives(stars []int) []int {
	var stack []int
	for i := range stars {
		curr := stars[i]
		if curr > 0 {
			stack = append(stack, curr)
			continue
		}

		survive := true
		for len(stack) > 0 && stack[len(stack)-1] > 0 {
			top := stack[len(stack)-1]
			if top == -curr {
				// 质量相当，都毁灭
				survive = false
				stack = stack[:len(stack)-1]
				break
			} else if top > -curr {
				// 向右的质量更大，向左的直接毁灭
				survive = false
				break
			} else {
				// 向左的质量更大，栈顶向右的直接毁灭
				stack = stack[:len(stack)-1]
			}
		}

		if survive {
			stack = append(stack, curr)
		}
	}

	return stack
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fields := strings.Split(
			strings.TrimRight(strings.TrimLeft(
				strings.TrimSpace(scanner.Text()), "["), "]"),
			",")
		arr := getIntArr(fields)
		res := Survives(arr)
		fmt.Println(res)
	}
}
