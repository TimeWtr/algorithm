// 题目：统计道路上的碰撞次数
//在一条无限长的公路上有 n 辆汽车正在行驶。汽车按从左到右的顺序按从 0 到 n - 1 编号，每辆车都在一个 独特的 位置。
//
//给你一个下标从 0 开始的字符串 directions ，长度为 n 。directions[i] 可以是 'L'、'R' 或 'S' 分别表示第 i 辆车是向 左 、向 右 或者 停留 在当前位置。每辆车移动时 速度相同 。
//
//碰撞次数可以按下述方式计算：
//
//当两辆移动方向 相反 的车相撞时，碰撞次数加 2 。
//当一辆移动的车和一辆静止的车相撞时，碰撞次数加 1 。
//碰撞发生后，涉及的车辆将无法继续移动并停留在碰撞位置。除此之外，汽车不能改变它们的状态或移动方向。
//
//返回在这条道路上发生的 碰撞总次数 。
//
//
//
//示例 1：
//
//输入：directions = "RLRSLL"
//输出：5
//解释：
//将会在道路上发生的碰撞列出如下：
//- 车 0 和车 1 会互相碰撞。由于它们按相反方向移动，碰撞数量变为 0 + 2 = 2 。
//- 车 2 和车 3 会互相碰撞。由于 3 是静止的，碰撞数量变为 2 + 1 = 3 。
//- 车 3 和车 4 会互相碰撞。由于 3 是静止的，碰撞数量变为 3 + 1 = 4 。
//- 车 4 和车 5 会互相碰撞。在车 4 和车 3 碰撞之后，车 4 会待在碰撞位置，接着和车 5 碰撞。碰撞数量变为 4 + 1 = 5 。
//因此，将会在道路上发生的碰撞总次数是 5 。
//示例 2：
//
//输入：directions = "LLRR"
//输出：0
//解释：
//不存在会发生碰撞的车辆。因此，将会在道路上发生的碰撞总次数是 0 。

// 题解：使用栈来解决，栈内存储的是向右或者停止的汽车，向左的汽车不断的与栈顶汽车
// 发生碰撞，碰撞后汽车停止会被压入栈顶

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// willCollide 是否发生碰撞，碰撞的条件是：
// 1. 如果栈顶为R：C必须是向左的或者是静止不动的
// 2. 如果栈顶是S：C必须是向左的L才能发生碰撞
func willCollide(top, c byte) bool {
	if top == 'R' && (c == 'L' || c == 'S') {
		return true
	}

	if top == 'S' && c == 'L' {
		return true
	}

	return false
}

// TODO 需要修改，解法有问题
func collisions(directions string) int {
	var stack []byte
	count := 0
	for i := range directions {
		c := directions[i]
		if c == 'R' {
			stack = append(stack, c)
			continue
		}

		for len(stack) > 0 && willCollide(stack[len(stack)-1], c) {
			top := stack[len(stack)-1]
			if top == 'R' && c == 'L' {
				// 两辆车碰撞
				count += 2
				stack = stack[:len(stack)-1]
				c = 'S'
			} else if top == 'R' && c == 'S' {
				// 一辆车静止
				count++
				stack = stack[:len(stack)-1]
			} else if top == 'S' && c == 'L' {
				count++
				c = 'S'
			} else {
				break
			}
		}
	}

	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		directions := strings.TrimSpace(scanner.Text())
		count := collisions(directions)
		fmt.Println(count)
	}
}
