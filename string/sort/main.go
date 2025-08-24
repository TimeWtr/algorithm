package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// 自定义比较函数，实现题目要求的排序规则
func wordCompare(a, b string) int {
	// 将单词转换为小写进行比较
	aLower := strings.ToLower(a)
	bLower := strings.ToLower(b)

	// 确定最小长度
	minLen := len(aLower)
	if len(bLower) < minLen {
		minLen = len(bLower)
	}

	// 逐个字符比较
	for i := 0; i < minLen; i++ {
		if aLower[i] < bLower[i] {
			return -1
		}
		if aLower[i] > bLower[i] {
			return 1
		}
	}

	// 前缀相同，短单词优先
	if len(aLower) != len(bLower) {
		if len(aLower) < len(bLower) {
			return -1
		}
		return 1
	}

	// 完全相等（不区分大小写）
	return 0
}

// 处理字符串排序和去重
func sortAndDeduplicate(input string) string {
	// 按空格分割单词
	words := strings.Fields(input)
	if len(words) == 0 {
		return ""
	}

	// 自定义排序
	sort.Slice(words, func(i, j int) bool {
		return wordCompare(words[i], words[j]) < 0
	})

	// 去重处理（使用栈）
	result := []string{words[0]}
	for i := 1; i < len(words); i++ {
		// 忽略大小写比较
		if strings.EqualFold(result[len(result)-1], words[i]) {
			continue // 重复单词，跳过
		}
		result = append(result, words[i])
	}

	return strings.Join(result, " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 处理多行输入（每行一组测试）
	for scanner.Scan() {
		input := scanner.Text()
		result := sortAndDeduplicate(input)
		fmt.Println(result)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "读取输入出错:", err)
	}
}
