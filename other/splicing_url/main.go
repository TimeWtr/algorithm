package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SplicingURL(str string) string {
	if str == "" {
		return ""
	}

	if str == "," {
		return "/"
	}

	arr := strings.SplitN(str, ",", 2)
	if len(arr) != 2 {
		return ""
	}

	// 都存在/
	if strings.HasSuffix(arr[0], "/") && strings.HasPrefix(arr[1], "/") {
		return fmt.Sprintf("%s%s", arr[0], arr[1][1:])
	}

	// 都不存在
	if !strings.HasSuffix(arr[0], "/") && !strings.HasPrefix(arr[1], "/") {
		return fmt.Sprintf("%s/%s", arr[0], arr[1])
	}

	// 只存在一个
	return fmt.Sprintf("%s%s", arr[0], arr[1])
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		url := SplicingURL(str)
		fmt.Println(url)
	}
}
