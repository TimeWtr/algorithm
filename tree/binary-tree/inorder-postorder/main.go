// 题目：二叉树的广度优先遍历
// 题目描述
//有一棵二叉树，每个节点由一个大写字母标识(最多26个节点）。
//
//现有两组字母，分别表示后序遍历（左孩子->右孩子->父节点）和中序遍历（左孩子->父节点->右孩子）的结果，请你输出层序遍历的结果。
//
//输入描述
//每个输入文件一行，第一个字符串表示后序遍历结果，第二个字符串表示中序遍历结果。（每串只包含大写字母）
//
//中间用单空格分隔。
//
//输出描述
//输出仅一行，表示层序遍历的结果，结尾换行。
//
//用例1
//输入
//CBEFDA CBAEDF
//输出
//ABDCEF
//说明
//二叉树为：
//
//     A
//    /   \
//  B     D
// /      /  \
//C      E    F
//
// 题解：
// 1. 构建二叉树
// 	  1.1 后序遍历最后一个元素就是根节点
//    1.2 根据根节点在中序遍历数组中找到对应的下标
//    1.3 中序遍历数组中跟节点左半部分为左子树，右半部分为右子树，分别计算长度
//    1.4 根据长度来找到后序遍历中的左右子树，然后递归构建二叉树
// 2. 广度优先遍历二叉树，使用Queue队列来统计

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

func buildBinaryTree(inorder, postorder []string) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}

	// 从中序遍历中找到根节点的位置
	rootValIndex := 0
	for i := range inorder {
		if inorder[i] == rootVal {
			rootValIndex = i
			break
		}
	}
	// 根据中序遍历找到左右子树
	leftTree := inorder[:rootValIndex]
	rightTree := inorder[rootValIndex+1:]

	// 根据左右子树的长度在后序遍历中找到对应的左右子树
	postLeftTree := postorder[:len(leftTree)]
	postRightTree := postorder[len(leftTree) : len(postorder)-1]

	// 构建树
	root.Left = buildBinaryTree(leftTree, postLeftTree)
	root.Right = buildBinaryTree(rightTree, postRightTree)
	return root
}

func levelTraverse(tree *TreeNode) string {
	if tree == nil {
		return ""
	}

	var builder strings.Builder
	queue := []*TreeNode{tree}
	for len(queue) > 0 {
		size := len(queue)
		for i := range queue {
			builder.WriteString(queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		queue = queue[size:]
	}

	return builder.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fields := strings.Fields(strings.TrimSpace(scanner.Text()))
		postorder := strings.Split(fields[0], "")
		inorder := strings.Split(fields[1], "")

		tree := buildBinaryTree(inorder, postorder)
		res := levelTraverse(tree)
		fmt.Println(res)
	}
}
