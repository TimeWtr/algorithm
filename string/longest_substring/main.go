//题目描述
//给定一个字符串，只包含字母和数字，按要求找出字符串中的最长（连续）子串的长度，字符串本身是其最长的子串，子串要求：
//
//1、 只包含1个字母(a~z, A~Z)，其余必须是数字；
//
//2、 字母可以在子串中的任意位置；
//
//如果找不到满足要求的子串，如全是字母或全是数字，则返回-1。
//
//输入描述
//字符串(只包含字母和数字)
//
//输出描述
//子串的长度
//
//用例
//输入	abC124ACb
//输出	4
//说明	满足条件的最长子串是C124或者124A，长度都是4
//输入	a5
//输出	2
//说明	字符串自身就是满足条件的子串，长度为2
//输入	aBB9
//输出	2
//说明	满足条件的子串为B9，长度为2
//输入	abcdef
//输出	-1
//说明	没有满足要求的子串，返回-1
//题目解析
//此题可以用滑动窗口求解。
//
//滑动窗口的左指针开始指向索引0，右指针也是从索引0开始不断向右移动，当右指针遇到字母时，则滑动窗口内部含字母量+1。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func handler(str string) int {
	if len(str) < 2 {
		return -1
	}

	left, letterCount := 0, 0
	maxLength := -1
	runes := []rune(str)
	for right := 0; right < len(runes); right++ {
		if unicode.IsLetter(runes[right]) {
			letterCount++
		}

		// 需要处理超过多余的字母字符
		for letterCount > 1 {
			if unicode.IsLetter(runes[left]) {
				letterCount--
			}
			left++
		}

		// 需要处理刚好为一个字母的统计比较
		if letterCount == 1 && right-left+1 >= 2 {
			if right-left+1 > maxLength {
				maxLength = right - left + 1
			}
		}
	}

	return maxLength
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := strings.TrimSpace(scanner.Text())
		length := handler(str)
		fmt.Println(length)
	}
}
