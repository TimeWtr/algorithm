package main

import "fmt"

type BinaryTree struct {
	// 节点值
	Val int
	// 左子树
	Left *BinaryTree
	// 右子树
	Right *BinaryTree
}

// 构造二叉搜索树
func buildBST(cur []int, lowIdx, highIdx int) *BinaryTree {
	if cur == nil {
		return nil
	}

	if lowIdx > highIdx {
		return nil
	}

	var midIdx int
	midIdx = lowIdx + (highIdx-lowIdx)/2
	root := &BinaryTree{}
	root.Left = buildBST(cur, lowIdx, midIdx-1)
	root.Right = buildBST(cur, midIdx+1, highIdx)
	root.Val = cur[midIdx]
	return root
}

func main() {
	cur := []int{1, 2, 3, 4, 5, 6, 7}
	root := buildBST(cur, 0, len(cur)-1)
	fmt.Println("BST: ", root)

	// 层序遍历打印所有结果
	var res []int
	queue := []*BinaryTree{root}
	for len(queue) > 0 {
		queueSize := len(queue)
		for i := 0; i < queueSize; i++ {
			res = append(res, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[queueSize:]
	}
	fmt.Println("遍历结果：", res)
}
