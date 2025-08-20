package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BinaryTree struct {
	Val         int
	Left, Right *BinaryTree
}

func levelPrint(pv, iv []int) []int {
	// 根据前序遍历和中序遍历构建出原始的二叉树
	originalTree := buildOriginTree(pv, iv)
	// 根据原是的二叉树构建求和二叉树
	sumTree := buildSumTree(originalTree)
	// 中序遍历列出最终的结果
	return inorder(sumTree)
}

func inorder(tree *BinaryTree) []int {
	var res []int
	if tree == nil {
		return res
	}

	var dfs func(node *BinaryTree)
	dfs = func(node *BinaryTree) {
		if node == nil {
			return
		}

		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}

	dfs(tree)
	return res
}

func buildSumTree(originalTree *BinaryTree) *BinaryTree {
	// 边界处理
	if originalTree == nil {
		return nil
	}

	var handler func(node *BinaryTree) (*BinaryTree, int)
	handler = func(node *BinaryTree) (*BinaryTree, int) {
		if node == nil {
			return nil, 0
		}

		leftNode, leftSum := handler(node.Left)
		rightNode, rightSum := handler(node.Right)
		newNode := &BinaryTree{
			Val:   leftSum + rightSum,
			Left:  leftNode,
			Right: rightNode,
		}
		return newNode, leftSum + rightSum + node.Val
	}

	sumTree, _ := handler(originalTree)
	return sumTree
}

func buildOriginTree(pv, iv []int) *BinaryTree {
	// 边界条件
	if len(pv) == 0 || len(iv) == 0 || len(pv) != len(iv) {
		return nil
	}

	// 取前序遍历第一个节点为根节点
	rootVal := pv[0]
	root := &BinaryTree{Val: rootVal}
	rootIndex := -1
	for i, v := range iv {
		if v == rootVal {
			rootIndex = i
			break
		}
	}

	// 边界处理
	if rootIndex == -1 {
		return nil
	}

	// 根据中序遍历根节点下标来计算左右子树的长度
	leftLength := rootIndex
	// 左子树中序遍历的长度等于前序遍历中左子树的长度，右子树同理
	root.Left = buildOriginTree(pv[1:leftLength+1], iv[:rootIndex])
	root.Right = buildOriginTree(pv[leftLength+1:], iv[rootIndex+1:])

	return root
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		var inorderVal []int
		for _, val := range strings.Split(strings.Trim(scanner.Text(), " "), " ") {
			iv, _ := strconv.Atoi(val)
			inorderVal = append(inorderVal, iv)
		}

		if !scanner.Scan() {
			break
		}
		var preorderVal []int
		for _, val := range strings.Split(strings.Trim(scanner.Text(), " "), " ") {
			pv, _ := strconv.Atoi(val)
			preorderVal = append(preorderVal, pv)
		}

		res := levelPrint(preorderVal, inorderVal)
		fmt.Println(res)
	}
}
