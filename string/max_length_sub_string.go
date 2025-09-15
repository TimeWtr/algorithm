// 题目：无重复字符的最长子串
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
//
//
//
//示例 1:
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
//提示：
//
//0 <= s.length <= 5 * 104
//s 由英文字母、数字、符号和空格组成

// 解法：使用滑动窗口+Map来统计字符

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	// 创建映射存储字符出现的最后位置
	// key: 字符，val: 下标位置
	m := make(map[byte]int)
	start, maxLen := 0, 0

	for i := 0; i < len(s); i++ {
		char := s[i]
		// 字符存在且字符出现字符在窗口内
		if idx, ok := m[char]; ok && idx >= start {
			// 这个字符已经是重复出现了，需要将前边的字符去掉
			start = idx + 1
		}

		// 更新当前字符的位置
		m[char] = i
		// 计算当前窗口的大小
		currLen := i - start + 1
		if maxLen < currLen {
			maxLen = currLen
		}
	}

	return maxLen
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		l := lengthOfLongestSubstring(line)
		fmt.Println(l)
	}
}
