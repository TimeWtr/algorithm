package main

import "fmt"

type BinaryTree struct {
	// 当前节点的值
	Val int
	// 左子树
	Left *BinaryTree
	// 右子树
	Right *BinaryTree
}

// 构建二叉搜索树
func buildBST(cur []int, lowIdx, highIdx int) *BinaryTree {
	if len(cur) == 0 || lowIdx > highIdx {
		return nil
	}

	bst := new(BinaryTree)
	midIdx := lowIdx + (highIdx-lowIdx)/2
	bst.Val = cur[midIdx]
	bst.Left = buildBST(cur, lowIdx, midIdx-1)
	bst.Right = buildBST(cur, midIdx+1, highIdx)
	return bst
}

func levelTraversal(tree *BinaryTree) []int {
	if tree == nil {
		return nil
	}

	var res []int
	queue := []*BinaryTree{tree}
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

	return res
}

// 计算路径和(递推思想实现)
func calculateSums(root *BinaryTree) []int {
	if root == nil {
		return nil
	}

	var results []int
	nodeQ := []*BinaryTree{root}
	sumsQ := []int{0}
	for len(nodeQ) > 0 {
		curNode := nodeQ[0]
		nodeQ = nodeQ[1:]
		curSum := sumsQ[0] + curNode.Val
		sumsQ = sumsQ[1:]
		if curNode.Left == nil && curNode.Right == nil {
			// 叶子节点
			results = append(results, curSum)
			continue
		}

		if curNode.Left != nil {
			nodeQ = append(nodeQ, curNode.Left)
			sumsQ = append(sumsQ, curSum)
		}

		if curNode.Right != nil {
			nodeQ = append(nodeQ, curNode.Right)
			sumsQ = append(sumsQ, curSum)
		}
	}

	return results
}

// 计算路径和(递归思想实现)
func calculateSumsRecursion(root *BinaryTree) []int {
	if root == nil {
		return nil
	}

	var sums []int
	var dfs func(node *BinaryTree, sum int)
	dfs = func(node *BinaryTree, sum int) {
		if node == nil {
			return
		}

		curSum := sum + node.Val
		if node.Left == nil && node.Right == nil {
			sums = append(sums, curSum)
			return
		}

		if node.Left != nil {
			dfs(node.Left, curSum)
		}

		if node.Right != nil {
			dfs(node.Right, curSum)
		}
	}

	dfs(root, 0)
	return sums
}

func main() {
	cur := []int{1, 2, 3, 4, 5, 6, 7}
	bst := buildBST(cur, 0, len(cur)-1)
	fmt.Println("构建二叉搜索树：", bst)
	res := levelTraversal(bst)
	fmt.Println("层级遍历二叉搜索树：", res)
	sums := calculateSums(bst)
	fmt.Println("路径和(递推)：", sums)
	sums1 := calculateSumsRecursion(bst)
	fmt.Println("路径和(递归)：", sums1)
}
