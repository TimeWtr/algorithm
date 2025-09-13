// 题目：二叉树的中序遍历
// 题目描述
//根据给定的二叉树结构描述字符串，输出该二叉树按照中序遍历结果字符串。中序遍历顺序为：左子树，根结点，右子树。
//
//输入描述
//由大小写字母、左右大括号、逗号组成的字符串:字母代表一个节点值，左右括号内包含该节点的子节点。
//
//左右子节点使用逗号分隔，逗号前为空则表示左子节点为空，没有逗号则表示右子节点为空。
//
//二叉树节点数最大不超过100。
//
//注:输入字符串格式是正确的，无需考虑格式错误的情况。
//
//输出描述
//输出一个字符串为二叉树中序遍历各节点值的拼接结果。
//
//用例1
//输入
//a{b{d,e{g,h{,i}}},c{f}}
//输出
//dbgehiafc
//
//
//
//
//
//
//题解
//思路：
//
//使用栈模拟递归处理二叉树节点的父子关系映射，构建出完整的二叉树结果。有一个小技巧，遍历过程节点没有右节点或子节点时也往栈中压入一个空节点。这样可以减少处理很多边界条件处理。
//进行中序遍历。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TreeNode struct {
	Val   string
	Left  *TreeNode
	Right *TreeNode
}

func buildBinaryTree(line string) *TreeNode {
	var (
		root  *TreeNode
		curr  *TreeNode
		stack []*TreeNode
	)
	isLeft := true

	for i := range line {
		char := line[i]
		switch char {
		case '{':
			if curr != nil {
				stack = append(stack, curr)
			}
			curr = nil
			isLeft = true
		case '}':
			if len(stack) > 0 {
				curr = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		case ',':
			isLeft = false
			curr = nil
		default:
			// 字母-需要处理节点
			curr = &TreeNode{Val: string(char)}
			if root == nil {
				root = curr
			} else if len(stack) > 0 {
				parent := stack[len(stack)-1]
				if isLeft {
					parent.Left = curr
				} else {
					parent.Right = curr
				}
			}
		}
	}

	return root
}

func inorder(tree *TreeNode) string {
	if tree == nil {
		return ""
	}
	return inorder(tree.Left) + tree.Val + inorder(tree.Right)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		tree := buildBinaryTree(line)
		res := inorder(tree)
		fmt.Println(res)
	}
}
