// 题目：考勤信息统计
// 题目描述
// 公司用一个字符串来表示员工的出勤信息

// absent：缺勤
// late：迟到
// leaveearly：早退
// present：正常上班
// 现需根据员工出勤信息，判断本次是否能获得出勤奖，能获得出勤奖的条件如下：

// 缺勤不超过一次；
// 没有连续的迟到/早退；
// 任意连续7次考勤，缺勤/迟到/早退不超过3次。
// 输入描述
// 用户的考勤数据字符串

// 记录条数 >= 1；
// 输入字符串长度 < 10000；
// 不存在非法输入；
// 如：

// 2
// present
// present absent present present leaveearly present absent

// 输出描述
// 根据考勤数据字符串，如果能得到考勤奖，输出”true”；否则输出”false”，
// 对于输入示例的结果应为：

// true false

// 示例1
// 输入

// 2
// present
// present present
// 1
// 2
// 3
// 输出

// true true
// 1
// 说明

// 无

// 示例2
// 输入

// 2
// present
// present absent present present leaveearly present absent
// 1
// 2
// 3
// 输出

// true false
// 1
//
// 解题思路：
// 1. 缺勤不超过一次：可以通过全局计数器一次遍历就可以完成
// 2. 没有连续的迟到/早退：通过一次遍历也可以完成
// 3. 任意连续7次考勤，缺勤/迟到/早退不超过3次：通过滑动窗口和计数器来实现

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkAttendance(arr []string) bool {
	// 判断缺勤是否超过一次
	absentCount := 0
	for _, v := range arr {
		if v == "absent" {
			absentCount++
			if absentCount > 1 {
				return false
			}
		}
	}

	// 判断是否有连续的迟到/早退
	for i := range len(arr) - 1 {
		curr, next := arr[i], arr[i+1]
		if (curr == "late" || curr == "leaveearly") && (next == "late" || next == "leaveearly") {
			return false
		}
	}

	// 滑动窗口算法来检查任意连续7次考勤，缺勤/迟到/早退不超过3次
	// 1. 处理考勤数小于7次的情况
	// 2. 超过7次需要先初始化第一个窗口，然后进行窗口滑动，窗口大小为7
	if len(arr) <= 7 {
		abnormalCount := 0
		for _, v := range arr {
			if v != "present" {
				abnormalCount++
			}
		}

		if abnormalCount > 3 {
			return false
		}

		return true
	}

	abnormalCount := 0
	// 初始化第一个窗口
	for i := range 7 {
		val := arr[i]
		if val != "present" {
			abnormalCount++
		}

		if abnormalCount > 3 {
			return false
		}
	}

	for i := 7; i < len(arr); i++ {
		// 去掉过期窗口的异常考勤数据
		if arr[i-7] != "present" {
			abnormalCount--
		}

		if arr[i] != "present" {
			abnormalCount++
			if abnormalCount > 3 {
				return false
			}
		}
	}

	return abnormalCount > 3
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		count := 0
		for count < n {
			if !scanner.Scan() {
				break
			}
			records := strings.Split(strings.TrimSpace(scanner.Text()), " ")
			res := checkAttendance(records)
			fmt.Println(res)
			count++
		}
	}
}
