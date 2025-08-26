// 题目描述
// 定义构造三叉搜索树规则如下：

// 每个节点都存有一个数，当插入一个新的数时，从根节点向下寻找，直到找到一个合适的空节点插入。查找的规则是：

// 如果数小于节点的数减去500，则将数插入节点的左子树
// 如果数大于节点的数加上500，则将数插入节点的右子树
// 否则，将数插入节点的中子树
// 给你一系列数，请按以上规则，按顺序将数插入树中，构建出一棵三叉搜索树，最后输出树的高度。
// 输入描述
// 第一行为一个数 N，表示有 N 个数，1 ≤ N ≤ 10000
// 第二行为 N 个空格分隔的整数，每个数的范围为[1,10000]

// 输出描述
// 输出树的高度（根节点的高度为1）

// 用例
// 输入	5
// 5000 2000 5000 8000 1800
// 输出	3
// 说明
// 最终构造出的树如下，高度为3：

// 输入	3
// 5000 4000 3000
// 输出	3
// 说明
// 最终构造出的树如下，高度为3：

// 输入	9
// 5000 2000 5000 8000 1800 7500 4500 1400 8100
// 输出	4
// 说明
// 最终构造出的树如下，高度为4：

// 解体思路：
// 1. 定义一个三叉树的结构，包括左子树、中子树、右子树和节点值
// 2. 根据规则插入对应节点值，插入前需要判断节点值是否为空，如果为空，直接插入
// 3. 如果节点值不为空，则需要插入子节点对应子树值，例如：左子节点值存在，则根据
// 当前节点找到该节点的左子节点作为插入节点。
// 4. 树的深度计算：树的高度为树的深度
// 	4.1 计算方式1: 从根节点开始，递归计算左右子树的高度，取较大值加1。
//  4.2 计算方式2: 在构建三叉树前声明一个最大深度变量，每次插入节点时如果需要新插入一层则
// 深度+1并比较最大深度，如果超过则更新最大深度。
// 当前实现基于方式2来进行。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const delta = 500

type TreeNode struct {
	// 节点值
	Val int
	// 左子节点
	Left *TreeNode
	// 中子节点
	Mid *TreeNode
	// 右子节点
	Right *TreeNode
}

// buildTreeWithHeight 构建三叉树并计算树的深度
func buildTreeWithHeight(arr []int) int {
	// 边界条件处理
	if len(arr) == 0 {
		return 0
	}

	root := &TreeNode{Val: arr[0]}
	maxDepth := 1
	for _, val := range arr[1:] {
		curr := root
		depth := 1
		for {
			depth++
			if val < curr.Val-delta {
				// 左子树
				if curr.Left == nil {
					curr.Left = &TreeNode{Val: val}
					if depth > maxDepth {
						maxDepth = depth
					}
					break
				}
				curr = curr.Left
			} else if val > curr.Val+delta {
				// 右子树
				if curr.Right == nil {
					curr.Right = &TreeNode{Val: val}
					if depth > maxDepth {
						maxDepth = depth
					}
					break
				}
				curr = curr.Right
			} else {
				// 中子树
				if curr.Mid == nil {
					curr.Mid = &TreeNode{Val: val}
					if depth > maxDepth {
						maxDepth = depth
					}
					break
				}
				curr = curr.Mid
			}
		}
	}

	return maxDepth
}

func handleArr(line string) []int {
	if line == "" {
		return nil
	}

	nums := strings.Split(strings.TrimSpace(line), " ")
	if len(nums) == 0 {
		return nil
	}

	res := make([]int, 0, len(nums))
	for _, val := range nums {
		n, _ := strconv.Atoi(val)
		res = append(res, n)
	}

	return res
}

// buildTree 构建三叉树
func buildTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	// 构建根节点
	root := &TreeNode{Val: arr[0]}
	for _, val := range arr[1:] {
		curr := root
		for {
			if val < curr.Val-delta {
				// 左子树
				if curr.Left == nil {
					curr.Left = &TreeNode{Val: val}
					break
				}
				curr = curr.Left
			} else if val > curr.Val+delta {
				// 右子树
				if curr.Right == nil {
					curr.Right = &TreeNode{Val: val}
					break
				}
				curr = curr.Right
			} else {
				// 中子树
				if curr.Mid == nil {
					curr.Mid = &TreeNode{Val: val}
					break
				}
				curr = curr.Mid
			}
		}
	}

	return root
}

// levelCalDepth 层级遍历计算树的深度
func levelCalDepth(tree *TreeNode) int {
	if tree == nil {
		return 0
	}

	nodeQ := []*TreeNode{tree}
	maxDepth := 0
	for len(nodeQ) > 0 {
		maxDepth++
		qSize := len(nodeQ)
		for i := 0; i < qSize; i++ {
			curr := nodeQ[i]
			if curr.Left != nil {
				nodeQ = append(nodeQ, curr.Left)
			}

			if curr.Mid != nil {
				nodeQ = append(nodeQ, curr.Mid)
			}

			if curr.Right != nil {
				nodeQ = append(nodeQ, curr.Right)
			}
		}
		nodeQ = nodeQ[qSize:]
	}

	return maxDepth
}

// recurDepth 递归计算最大深度
func recurDepth(height int, node *TreeNode) int {
	if node == nil {
		return height
	}

	return max(max(recurDepth(height, node.Left), recurDepth(height, node.Mid)),
		recurDepth(height, node.Right)) + 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		n, _ := strconv.Atoi(line)

		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
		arr := handleArr(line)
		if len(arr) != n {
			fmt.Println("length not equal")
			break
		}

		maxDepth := buildTreeWithHeight(arr)
		fmt.Println("max depth: ", maxDepth)

		// 先构建三叉树
		tree := buildTree(arr)
		// 广度优先计算最大深度
		fmt.Println("tree height: ", levelCalDepth(tree))
		// 深度优先递归计算最大深度
		maxDepth = recurDepth(0, tree)
		fmt.Println("tree depth: ", maxDepth)
	}
}
