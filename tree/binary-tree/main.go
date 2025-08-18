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

// calculateSums 计算所有的路径和(层级遍历，递推的思想实现)
func calculateSums(tree *BinaryTree) []int {
	if tree == nil {
		return nil
	}

	var res []int
	nodeQ := []*BinaryTree{tree}
	sumsQ := []int{0}
	for len(nodeQ) > 0 {
		curNode := nodeQ[0]
		nodeQ = nodeQ[1:]
		curSum := curNode.Val + sumsQ[0]
		sumsQ = sumsQ[1:]
		if curNode.Left == nil && curNode.Right == nil {
			// 没有子节点，当前节点为叶子节点
			res = append(res, curSum)
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

	return res
}

// calculateSumsRecursion 计算所有的路径和(深度优先，递归代码实现递推思想)
func calculateSumsRecursion(tree *BinaryTree) []int {
	if tree == nil {
		return nil
	}

	var res []int
	var dfs func(node *BinaryTree, sum int)
	dfs = func(node *BinaryTree, sum int) {
		if node == nil {
			return
		}

		curSum := node.Val + sum
		if node.Left == nil && node.Right == nil {
			// 叶子节点
			res = append(res, curSum)
			return
		}

		if node.Left != nil {
			dfs(node.Left, curSum)
		}

		if node.Right != nil {
			dfs(node.Right, curSum)
		}
	}

	dfs(tree, 0)
	return res
}

// sumsRecursion 计算所有路径和(递归代码实现递归思想)
func sumsRecursion(tree *BinaryTree) []int {
	if tree == nil {
		return nil
	}

	sums := sumsRecursion(tree.Left)
	rightSums := sumsRecursion(tree.Right)
	// 汇总所有路径
	sums = append(sums, rightSums...)
	// 加上当前节点的值
	for i := 0; i < len(sums); i++ {
		sums[i] = sums[i] + tree.Val
	}

	if len(sums) == 0 {
		// 没有子路径，那就只加上当前的节点值
		sums = append(sums, tree.Val)
	}

	return sums
}

// collectSumsBottomUp 计算路径和，以自底向上的思想来实现，递归处理，左右子树
// 分别累加当前节点值写入到sums中，避免合并后再累加，防止出现大的切片合并操作
func collectSumsBottomUp(tree *BinaryTree) []int {
	if tree == nil {
		return nil
	}

	var sums []int
	for _, leftSum := range collectSumsBottomUp(tree.Left) {
		sums = append(sums, leftSum+tree.Val)
	}

	for _, rightSum := range collectSumsBottomUp(tree.Right) {
		sums = append(sums, rightSum+tree.Val)
	}

	if len(sums) == 0 {
		// 没有左右子树，是一个叶子节点
		sums = append(sums, tree.Val)
	}

	return sums
}

func main() {
	cur := []int{1, 2, 3, 4, 5, 6, 7}
	bst := buildBST(cur, 0, len(cur)-1)
	fmt.Println("构建二叉搜索树：", bst)
	res := levelTraversal(bst)
	fmt.Println("层级遍历二叉搜索树：", res)
	sums := calculateSums(bst)
	fmt.Println("路径和(calculateSums)：", sums)
	sums1 := calculateSumsRecursion(bst)
	fmt.Println("路径和(calculateSumsRecursion)：", sums1)
	sums2 := sumsRecursion(bst)
	fmt.Println("路径和(sumsRecursion): ", sums2)
	sums3 := collectSumsBottomUp(bst)
	fmt.Println("路径和(collectSumsBottomUp): ", sums3)
}
