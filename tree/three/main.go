package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const delta = 500

type TernaryTree struct {
	Val              int
	Left, Mid, Right *TernaryTree
}

func buildTernaryTree(n int, arr []int) int {
	if n == 0 || len(arr) == 0 {
		return 0
	}

	// 构建根节点
	root := &TernaryTree{Val: arr[0]}
	maxDepth := 1
	for _, val := range arr[1:] {
		depth := 1
		current := root
		for {
			depth++
			if val < current.Val-delta {
				// 左子节点
				if current.Left == nil {
					current.Left = &TernaryTree{Val: val}
					if depth > maxDepth {
						maxDepth = depth
					}
					break
				}
				current = current.Left
			} else if val > current.Val+delta {
				if current.Right == nil {
					current.Right = &TernaryTree{Val: val}
					if depth > maxDepth {
						maxDepth = depth
					}
					break
				}
				current = current.Right
			} else {
				if current.Mid == nil {
					current.Mid = &TernaryTree{Val: val}
					if depth > maxDepth {
						maxDepth = depth
					}
					break
				}
				current = current.Mid
			}
		}
	}

	return maxDepth
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		var num int
		var arr []int
		if !scanner.Scan() {
			break
		}
		num, _ = strconv.Atoi(scanner.Text())

		if !scanner.Scan() {
			break
		}
		arrStr := strings.Split(strings.Trim(scanner.Text(), " "), " ")
		for _, str := range arrStr {
			num, _ = strconv.Atoi(str)
			arr = append(arr, num)
		}

		depth := buildTernaryTree(num, arr)
		fmt.Println("tree depth: ", depth)
	}
}
