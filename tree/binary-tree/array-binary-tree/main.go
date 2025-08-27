//题目：数组二叉树
//题目描述
//二叉树也可以用数组来存储，给定一个数组，树的根节点的值存储在下标1，对于存储在下标N的节点，它的左子节点和右子节点分别存储在下标2*N和2*N+1，并且我们用值-1代表一个节点为空。
//
//给定一个数组存储的二叉树，试求从根节点到最小的叶子节点的路径，路径由节点的值组成。
//
//输入描述
//输入一行为数组的内容，数组的每个元素都是正整数，元素间用空格分隔。
//
//注意第一个元素即为根节点的值，即数组的第N个元素对应下标N，下标0在树的表示中没有使用，所以我们省略了。
//
//输入的树最多为7层。
//
//输出描述
//输出从根节点到最小叶子节点的路径上，各个节点的值，由空格分隔，用例保证最小叶子节点只有一个。
//用例
//输入	3 5 7 -1 -1 2 4
//输出	3 7 2
//说明	最小叶子节点的路径为3 7 2。
//输入	5 9 8 -1 -1 7 -1 -1 -1 -1 -1 6
//输出	5 8 7 6
//说明	最小叶子节点的路径为5 8 7 6，注意数组仅存储至最后一个非空节点，故不包含节点“7”右子节点的-1。

//解题思路：
// 1. 数组根据下标已经可以推算出整棵树，所以不需要再构建树，因为构建需要额外的空间和时间
// 2. 数组下标为0的元素即为根节点，左子节点下标：2i+1，右子节点下标：2i+2，父节点的下标：(i-1)/2
// 左子节点求父节点下标直接得到的是整数，右子节点下标推算出来的父节点下标为浮点数，不过go直接会直接将
// 其转换为正整数，所以(i-1)/2可以直接推算出父节点下标。
// 解题方案：
// 1. 自顶向下：从根节点向叶子节点递归，找到所有的根基点到叶子节点的路径，对比最小叶子节点，然后再找出
// 对应的路径。
// 2. 自底向上：先找到最小的叶子节点，然后回溯找父节点，直至跟节点。

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// findPathTopDown 自顶向下，先找到最小节点的下标，然后通过回溯+反转切片获取最后的结果
func findPathTopDown(arr []int) string {
	if len(arr) == 0 {
		return ""
	}

	// 定义最小的叶子节点值和下标
	l := len(arr)
	minLeafVal := math.MaxInt32
	minLeafIndex := -1

	for i, val := range arr {
		// 跳过空节点(-1)
		if val == -1 {
			continue
		}

		// 找到左右节点的下标
		leftIndex := 2*i + 1
		rightIndex := 2*i + 2
		isleaf := true
		if leftIndex < l && arr[leftIndex] != -1 {
			isleaf = false
		}

		if rightIndex < l && arr[rightIndex] != -1 {
			isleaf = false
		}

		// 处理叶子节点
		if isleaf && arr[i] < minLeafVal {
			minLeafVal = arr[i]
			minLeafIndex = i
		}
	}

	// 根据最小的叶子节点下标推算父节点
	var path []string
	currIndex := minLeafIndex
	for currIndex >= 0 {
		path = append(path, strconv.Itoa(arr[currIndex]))
		if currIndex == 0 {
			break
		}
		parentIndex := (currIndex - 1) / 2
		currIndex = parentIndex
	}

	// 对路径进行反转
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return strings.Join(path, " ")
}

// findPathDFS 深度优先算法计算最小叶子节点的完整路径
func findPathDFS(arr []int) string {
	if len(arr) == 0 {
		return ""
	}

	// 定义递归闭包函数
	l := len(arr)
	minLeafVal := math.MaxInt32
	var minPath []string
	var dfs func(index int, path []string)
	dfs = func(index int, path []string) {
		// 边界条件：下标越界/空节点
		if index >= len(arr) || arr[index] == -1 {
			return
		}

		// 拷贝路径
		currentPath := make([]string, len(path))
		copy(currentPath, path)
		currentPath = append(currentPath, strconv.Itoa(arr[index]))

		// 处理左右子树
		leftIndex := 2*index + 1
		rightIndex := 2*index + 2
		isLeaf := true
		if leftIndex < l && arr[leftIndex] != -1 {
			isLeaf = false
			dfs(leftIndex, currentPath)
		}

		if rightIndex < l && arr[rightIndex] != -1 {
			isLeaf = false
			dfs(rightIndex, currentPath)
		}

		// 处理叶子节点
		if isLeaf && arr[index] < minLeafVal {
			minLeafVal = arr[index]
			minPath = currentPath
		}
	}

	dfs(0, []string{})

	return strings.Join(minPath, " ")
}

func handleArr(strArr []string) []int {
	if len(strArr) == 0 {
		return nil
	}

	res := make([]int, 0, len(strArr))
	for _, str := range strArr {
		val, _ := strconv.Atoi(str)
		res = append(res, val)
	}

	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strArr := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		arr := handleArr(strArr)
		if arr == nil {
			break
		}
		path := findPathTopDown(arr)
		fmt.Println(path)
		path1 := findPathDFS(arr)
		fmt.Println(path1)
	}
}
