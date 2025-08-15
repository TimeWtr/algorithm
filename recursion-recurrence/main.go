package main

import "fmt"

func main() {
	cur := []int{100, 200, 300, 400, 500, 600, 700, 800, 900}
	res := recurrence(cur)
	fmt.Println("递推结果：", res)

	res = recursion(cur, 0)
	fmt.Println("递归结果：", res)

	res = recursionWith(cur, 0)
	fmt.Println("递归无害补偿结果：", res)
}

// 递推实现
func recurrence(cur []int) int {
	var res int
	for i := 1; i < len(cur); i++ {
		res += cur[i]
	}

	return res
}

// 递归实现
func recursion(cur []int, level int) int {
	if level == len(cur)-1 {
		return cur[level]
	}

	return cur[level] + recursion(cur, level+1)
}

// 递归越界无害补偿实现
func recursionWith(cur []int, level int) int {
	// 越界补偿
	if level == len(cur) {
		return 0
	}

	return cur[level] + recursionWith(cur, level+1)
}
