package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func encrypt(words []string) string {
	if len(words) == 0 {
		return ""
	}

	// 循环判断每一个单词，判断是否存在元音字符
	for i, word := range words {
		exist := false
		for _, char := range word {
			if isVeo(char) {
				exist = true
				break
			}
		}

		if exist {
			// 存在需要加密处理
			words[i] = replaceVeo(word)
		} else {
			words[i] = replaceFirstLast(word)
		}

	}

	return strings.Join(words, " ")
}

func replaceVeo(word string) string {
	var builder strings.Builder
	for _, char := range word {
		if isVeo(char) {
			builder.WriteRune('*')
		} else {
			builder.WriteRune(char)
		}
	}

	return builder.String()
}

func replaceFirstLast(word string) string {
	chars := []rune(word)
	chars[0], chars[len(chars)-1] = chars[len(chars)-1], chars[0]
	return string(chars)
}

func isVeo(char rune) bool {
	char = unicode.ToLower(char)
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}

		wordsStr := strings.TrimSpace(scanner.Text())
		words := strings.Split(wordsStr, " ")
		res := encrypt(words)
		fmt.Println(res)
	}
}
