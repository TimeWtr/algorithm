// 题目：敏感字段加密
// 题目描述
// 给定一个由多个命令字组成的命令字符串：

// 1、字符串长度小于等于127字节，只包含大小写字母，数字，下划线和偶数个双引号；
// 2、命令字之间以一个或多个下划线_进行分割；
// 3、可以通过两个双引号””来标识包含下划线_的命令字或空命令字（仅包含两个双引号的命令字），双引号不会在命令字内部出现；

// 请对指定索引的敏感字段进行加密，替换为******（6个*），并删除命令字前后多余的下划线_。

// 如果无法找到指定索引的命令字，输出字符串ERROR。

// 输入描述
// 输入为两行，第一行为命令字索引K（从0开始），第二行为命令字符串S。

// 输出描述
// 输出处理后的命令字符串，如果无法找到指定索引的命令字，输出字符串ERROR

// 示例1
// 输入

// 1
// password__a12345678_timeout_100
// 1
// 2
// 输出

// password_******_timeout_100
// 1
// 说明

// 示例2
// 输入

// 2
// aaa_password_"a12_45678"_timeout__100_""_
// 1
// 2
// 输出

// aaa_password_******_timeout_100_""
// 1
// 说明

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func encryptSensitiveField(k int, cmd string) string {
	// 处理边界条件
	if cmd == "" {
		return "ERROR"
	}

	currToken := ""
	tokens := make([]string, 0, len(cmd))
	isQuotes := false
	for i := 0; i < len(cmd); i++ {
		char := cmd[i]
		if char == '"' {
			if isQuotes {
				// 这个是结束的双引号
				currToken += string(char)
				tokens = append(tokens, currToken)
				currToken = ""
				isQuotes = false
			} else {
				// 这个是开始的双引号
			}
		} else if char == '_' {
			if isQuotes {
				// 已经在引号之内了
				currToken += string(char)
			} else {
				// 不在引号内，需要将多余的下划线给去掉
				if currToken != "" {
					tokens = append(tokens, currToken)
					currToken = ""
				}
				    
				// 跳过连续的下划线
				for i < len(cmd)-1 && cmd[i+1] == '_' {
					i++
				}
			}
		} else {
			currToken += string(char)
		}
	}
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
		cmd := strings.TrimSpace(scanner.Text())
	}
}
