// 题目：字符串的统计与重排
//题目描述
// 给出一个仅包含字母的字符串，不包含空格，统计字符串中各个字母（区分大小写）出现的次数，

// 并按照字母出现次数从大到小的顺序。输出各个字母及其出现次数。

// 如果次数相同，按照自然顺序进行排序，且小写字母在大写字母之前。

// 输入描述
// 输入一行，为一个仅包含字母的字符串。

// 输出描述
// 按照字母出现次数从大到小的顺序输出各个字母和字母次数，用英文分号分隔，注意末尾的分号；

// 字母和次数间用英文冒号分隔。

// 示例1
// 输入

// xyxyXX
// 1
// 输出

// x:2;y:2;X:2;
// 1
// 说明

// 无

// 示例2
// 输入

// abababb
// 1
// 输出

// b:4;a:3;
//
// 解题思路：
// 1. 遍历整个字符串的字符，统计每一个字符出现的次数，区分大小写
// 2. 生成字符数组，根据ASCII进行进行排序
// 3. 输出最终的统计结果（字符串）

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type StrEntry struct {
	// 原始字符
	char rune
	// 出现的次数
	count int
}

func sortString(str string) string {
	if len(str) == 0 {
		return ""
	}

	// key: 字符, val: 出现的次数
	counterSet := map[rune]int{}
	for _, v := range str {
		counterSet[v]++
	}

	// 字符统计的排序数组
	sortedSet := make([]StrEntry, 0, len(counterSet))
	for k, v := range counterSet {
		sortedSet = append(sortedSet, StrEntry{
			char:  k,
			count: v,
		})
	}
	sort.SliceStable(sortedSet, func(i, j int) bool {
		// 先比较次数
		if sortedSet[i].count != sortedSet[j].count {
			return sortedSet[i].count > sortedSet[j].count
		}

		charI := sortedSet[i].char
		charJ := sortedSet[j].char
		isLowerI := charI >= 'a' && charI <= 'z'
		isLowerJ := charJ >= 'a' && charJ <= 'z'
		isUpperI := charI >= 'A' && charI <= 'Z'
		isUpperJ := charJ >= 'A' && charJ <= 'Z'

		// 一个为小写一个为大写字母,ASCII中大写字母小于小写字母
		if isLowerI && isUpperJ {
			return true
		}

		if isUpperI && isLowerJ {
			return false
		}

		// 都是大写/小写字母
		return charI < charJ
	})

	// 构造结果
	var res strings.Builder
	for _, entry := range sortedSet {
		res.WriteString(string(entry.char) + ":")
		counter, _ := counterSet[entry.char]
		res.WriteString(strconv.Itoa(counter))
		res.WriteString(";")
	}

	return res.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := strings.TrimSpace(scanner.Text())
		res := sortString(str)
		fmt.Println(res)
	}
}
