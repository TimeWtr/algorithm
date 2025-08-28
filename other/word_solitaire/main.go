// 题目：单词接龙
// 题目描述
// 单词接龙的规则是：

// 可用于接龙的单词首字母必须要前一个单词的尾字母相同；

// 当存在多个首字母相同的单词时，取长度最长的单词，如果长度也相等，则取字典序最小的单词；已经参与接龙的单词不能重复使用。

// 现给定一组全部由小写字母组成单词数组，并指定其中的一个单词作为起始单词，进行单词接龙，

// 请输出最长的单词串，单词串是单词拼接而成，中间没有空格。

// 输入描述
// 输入的第一行为一个非负整数，表示起始单词在数组中的索引K，0 <= K < N ；

// 输入的第二行为一个非负整数，表示单词的个数N；

// 接下来的N行，分别表示单词数组中的单词。

// 备注：

// 单词个数N的取值范围为[1, 20]；
// 单个单词的长度的取值范围为[1, 30]；
// 输出描述
// 输出一个字符串，表示最终拼接的单词串。

// 示例1
// 输入

// 0
// 6
// word
// dd
// da
// dc
// dword
// d
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 输出

// worddwordda
// 1
// 说明

// 先确定起始单词word，再接以d开头的且长度最长的单词dword，剩余以d开头且长度最长的有dd、da、dc，则取字典序最小的da，所以最后输出worddwordda。

// 示例2
// 输入

// 4
// 6
// word
// dd
// da
// dc
// dword
// d
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 输出

// dwordda
// 1
// 说明

// 先确定起始单词dword，剩余以d开头且长度最长的有dd、da、dc，则取字典序最小的da，所以最后输出dwordda。

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solitarie(k int, words []string) string {
	if len(words) == 0 {
		return ""
	}

	// 先找到接龙的首个单词
	firstWord := words[k]
	res := firstWord
	charSets := getCharets(k, words)
	// 已经被使用的单词的集合，key是单词
	usedSets := make(map[string]struct{})
	// 单词尾字符
	currSuffixChar := firstWord[len(firstWord)-1]
	for {
		// 候选单词数组
		candidates := charSets[currSuffixChar]
		found := false

		for _, word := range candidates {
			_, ok := usedSets[word]
			if ok {
				// 跳过已经使用过的单词
				continue
			}
			res += word
			found = true
			// 更新尾字符
			currSuffixChar = word[len(word)-1]
			usedSets[word] = struct{}{}
			break
		}

		if !found {
			break
		}
	}

	return res
}

// getCharets 根据单词首字母创建单词映射集合
// key: 单词的开头字符，val：所有以固定字符开头的单词数组
func getCharets(k int, words []string) map[byte][]string {
	charSets := make(map[byte][]string)
	for i, word := range words {
		// 跳过第一个单词
		if i == k {
			continue
		}
		char := []byte(word)[0]
		set, ok := charSets[char]
		if !ok {
			set = []string{}
		}
		set = append(set, word)
		sort.SliceStable(set, func(i, j int) bool {
			// 先比较单词长度
			if len(set[i]) != len(set[j]) {
				return len(set[i]) > len(set[j])
			}

			// 比较字典序
			return set[i] < set[j]
		})
		charSets[char] = set
	}

	return charSets
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		k, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		if !scanner.Scan() {
			break
		}
		n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		words := make([]string, 0, n)
		count := 0
		for count < n {
			if !scanner.Scan() {
				break
			}
			words = append(words, strings.TrimSpace(scanner.Text()))
			count++
		}
		res := solitarie(k, words)
		fmt.Println(res)
	}
}
