package main

import (
	"fmt"
)

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

// calculateDivisorPath 统计所有路径和中结果是指定数公约数的路径(递推的思想实现)
// 节点队列、路径队列和路径和队列，一一对应
func calculateDivisorPath(tree *BinaryTree, divisor int) [][]int {
	// 边界条件
	if tree == nil || divisor < 0 {
		return nil
	}

	var res [][]int
	nodeQ := []*BinaryTree{tree}
	paths := [][]int{{}}
	sumsQ := []int{0}
	for len(nodeQ) > 0 {
		curNode := nodeQ[0]
		curSum := curNode.Val + sumsQ[0]
		path := paths[0]
		nodeQ = nodeQ[1:]
		sumsQ = sumsQ[1:]
		paths = paths[1:]

		path = append(path, curNode.Val)
		if curNode.Left == nil && curNode.Right == nil {
			// 没有左右子树，是叶子节点，判断是否是约数
			if divisor%curSum == 0 {
				// 是公约数
				res = append(res, path)
				continue
			}
		}

		if curNode.Left != nil {
			nodeQ = append(nodeQ, curNode.Left)
			sumsQ = append(sumsQ, curSum)
			paths = append(paths, path)
		}

		if curNode.Right != nil {
			nodeQ = append(nodeQ, curNode.Right)
			sumsQ = append(sumsQ, curSum)
			paths = append(paths, path)
		}
	}

	return res
}

// calculateDivisorPathRecursion 统计所有路径和中结果是指定数公约数的路径(递归思想的实现)
func calculateDivisorPathRecursion(tree *BinaryTree, divisor int) [][]int {
	if tree == nil || divisor < 0 {
		return nil
	}

	var res [][]int
	var dfs func(node *BinaryTree, divisor, sum int, path []int) [][]int
	dfs = func(node *BinaryTree, divisor, sum int, path []int) [][]int {
		if node == nil {
			return nil
		}

		curSum := node.Val + sum
		path = append(path, node.Val)
		// 终止条件
		if node.Left == nil && node.Right == nil {
			// 叶子节点
			if divisor%curSum == 0 {
				// 是公约数
				res = append(res, path)
			}

			return res
		}

		if node.Left != nil {
			dfs(node.Left, divisor, curSum, path)
		}

		if node.Right != nil {
			dfs(node.Right, divisor, curSum, path)
		}

		return res
	}

	dfs(tree, divisor, 0, []int{})
	return res
}

// calculateLeafDivisorPath 计算所有的路径，这里是路径中叶子节点是能够被公约数整除的
// 所有路径(层级遍历递推)，自顶向下
func calculateLeafDivisorPath(tree *BinaryTree, divisor int) [][]int {
	if tree == nil || divisor < 0 {
		return nil
	}

	var res [][]int
	nodeQ := []*BinaryTree{tree}
	paths := [][]int{{}}
	for len(nodeQ) > 0 {
		curNode := nodeQ[0]
		path := paths[0]
		nodeQ = nodeQ[1:]
		paths = paths[1:]
		path = append(path, curNode.Val)
		if curNode.Left == nil && curNode.Right == nil {
			// 叶子节点
			if divisor%curNode.Val == 0 {
				// 是叶子节点的公约数
				res = append(res, path)
				continue
			}
		}

		if curNode.Left != nil {
			nodeQ = append(nodeQ, curNode.Left)
			paths = append(paths, path)
		}

		if curNode.Right != nil {
			nodeQ = append(nodeQ, curNode.Right)
			paths = append(paths, path)
		}
	}

	return res
}

// calculateLeafDivisorPathRecursion 计算所有的路径，这里是路径中叶子节点是
// 能够被公约数整除的所有路径(闭包+递归)，自顶向下
func calculateLeafDivisorPathRecursion(tree *BinaryTree, divisor int) [][]int {
	if tree == nil || divisor < 0 {
		return nil
	}

	var res [][]int
	var dfs func(node *BinaryTree, divisor int, path []int)
	dfs = func(node *BinaryTree, divisor int, path []int) {
		if node == nil {
			return
		}

		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil {
			// 叶子节点
			if divisor%node.Val == 0 {
				res = append(res, path)
			}
			return
		}

		if node.Left != nil {
			dfs(node.Left, divisor, path)
		}

		if node.Right != nil {
			dfs(node.Right, divisor, path)
		}
	}

	dfs(tree, divisor, []int{})
	return res
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
	paths := calculateDivisorPath(bst, 63)
	fmt.Println("公约数的路径：", paths)
	paths1 := calculateDivisorPathRecursion(bst, 63)
	fmt.Println("公约数的路径：", paths1)
	paths2 := calculateLeafDivisorPath(bst, 21)
	fmt.Println("公约数的路径：", paths2)
	paths3 := calculateLeafDivisorPathRecursion(bst, 21)
	fmt.Println("公约数的路径：", paths3)
}
