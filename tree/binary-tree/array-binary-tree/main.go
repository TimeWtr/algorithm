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

// findPathTopDown 自顶向下递归查找最小叶子节点路径
func findPathTopDown(arr []int) string {
	if arr == nil {
		return ""
	}

	// 找出最小叶子节点的下标
	minLeafVal := math.MaxInt32
	minLeafIndex := -1
	length := len(arr)
	for i := 0; i < len(arr); i++ {
		// 跳过空节点
		if arr[i] == -1 {
			continue
		}

		isLeaf := true
		leftIndex := 2*i + 1
		rightIndex := 2*i + 2
		if leftIndex < length && arr[leftIndex] != -1 {
			isLeaf = false
		}

		if rightIndex < length && arr[rightIndex] != -1 {
			isLeaf = false
		}

		// 处理叶子节点，条件：1. 是叶子节点，2. 当前叶子节点小于最小叶子节点值
		if isLeaf && arr[i] < minLeafVal {
			minLeafVal = arr[i]
			minLeafIndex = i
		}
	}

	// 根据最小叶子节点下标和值回溯找到对应的路径
	var path []int
	curIndex := minLeafIndex
	for curIndex >= 0 {
		path = append(path, arr[curIndex])
		if curIndex == 0 {
			break
		}
		parentIndex := (curIndex - 1) / 2
		curIndex = parentIndex
	}

	// 先将数组路径反转
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	var res string
	for i := 0; i < len(path); i++ {
		res += strconv.Itoa(path[i])
		if i < len(path)-1 {
			res += " "
		}
	}

	return res
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
	}
}
