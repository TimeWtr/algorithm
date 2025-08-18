package main

import "fmt"

// 递归思想来实现计数次数统计
func countsRecursion(n, count int) int {
	// 边界条件
	if n == 0 {
		return 0
	}

	if n%2 == 0 {
		// 偶数
		n = n / 2
	} else {
		n = n*3 + 1
	}
	count++
	// 终止条件
	if n == 1 {
		return count
	}
	return countsRecursion(n, count)
}

// 普通递归(闭包形式)实现
func countRecurrence(n int) int {
	// 边界条件
	if n == 0 {
		return 0
	}

	var count int
	var dfs func(n int)
	dfs = func(n int) {
		if n == 0 {
			return
		}

		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
		count++
		if n == 1 {
			return
		}
		dfs(n)
	}
	dfs(n)
	return count
}

// tailRecursion 尾递归(递归优化)
func tailRecursion(n, count int) int {
	if n == 1 {
		return count
	}

	if n%2 == 0 {
		return tailRecursion(n/2, count+1)
	}
	return tailRecursion(n*3+1, count+1)
}

// Iter 迭代方式来实现计算
func Iter(n int) int {
	if n == 0 {
		return 0
	}

	var count int
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
		count++
	}

	return count
}

func main() {
	count := countsRecursion(10, 0)
	fmt.Println("总计次数为：", count)
	count1 := countRecurrence(10)
	fmt.Println("总计次数为：", count1)
	count2 := countsRecursion(10000, 0)
	fmt.Println("总计次数为：", count2)
	count3 := countRecurrence(10000)
	fmt.Println("总计次数为：", count3)
	count4 := Iter(10)
	fmt.Println("总计次数为：", count4)
	count5 := tailRecursion(10, 0)
	fmt.Println("总计次数为: ", count5)
}
