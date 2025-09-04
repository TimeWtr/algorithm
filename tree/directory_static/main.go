// 题目：文件目录大小
// 题目描述
// 一个文件目录的数据格式为：目录id，本目录中文件大小，(子目录id列表）。

// 其中目录id全局唯一，取值范围[1, 200]，本目录中文件大小范围[1, 1000]，子目录id列表个数[0,10]例如 : 1 20 (2,3) 表示目录1中文件总大小是20，有两个子目录，id分别是2和3

// 现在输入一个文件系统中所有目录信息，以及待查询的目录 id ，返回这个目录和及该目录所有子目录的大小之和。

// 输入描述
// 第一行为两个数字M，N，分别表示目录的个数和待查询的目录id,

// 1 ≤ M ≤ 100
// 1 ≤ N ≤ 200
// 接下来M行，每行为1个目录的数据：

// 目录id 本目录中文件大小 (子目录id列表)

// 子目录列表中的子目录id以逗号分隔。

// 输出描述
// 待查询目录及其子目录的大小之和

// 用例1
// 输入
// 3 1
// 3 15 ()
// 1 20 (2)
// 2 10 (3)
// 输出
// 45
// 说明
//  目录1大小为20，包含一个子目录2 (大小为10)，子目录2包含一个子目录3(大小为15)，总的大小为20+10+15=45

// 用例2
// 输入
// 4 2
// 4 20 ()
// 5 30 ()
// 2 10 (4,5)
// 1 40 ()
// 输出
// 60
// 说明
// 目录2包含2个子目录4和5，总的大小为10+20+30 = 60
//
// 解题思路：这是一个树型结构，子节点在()内，可以先存储每一个目录节点的属性
// 再创建当前节点与子节点的映射关系，根据节点属性和映射关系显式的构建一颗树，
// 找到对应的节点，然后根据DFS算法计算目录的总大小。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DirectoryNode struct {
	// 目录ID
	ID int
	// 目录大小
	Size int64
	// 目录的子节点
	Children []*DirectoryNode
}

func buildTreeMap(arr []string) map[int]*DirectoryNode {
	// 创建node属性和节点与子节点的映射关系
	nodes := make(map[int]*DirectoryNode)
	// key: 目录ID，val：所有子目录的ID
	pendingRelationship := make(map[int][]int)
	for _, line := range arr {
		fields := strings.Fields(line)
		id, _ := strconv.Atoi(fields[0])
		size, _ := strconv.ParseInt(fields[1], 10, 64)
		// 构建节点属性
		nodes[id] = &DirectoryNode{
			ID:   id,
			Size: size,
		}

		// 处理目录节点与子节点之间的映射关系
		part := fields[2]
		if part != "()" {
			part = strings.TrimSuffix(strings.TrimPrefix(part, "("), ")")
			children := strings.Split(part, ",")
			for _, child := range children {
				intChild, _ := strconv.Atoi(child)
				pendingRelationship[id] = append(pendingRelationship[id], intChild)
			}
		}
	}

	// 构建完整的树结构
	for id, children := range pendingRelationship {
		parentNode := nodes[id]
		for _, child := range children {
			parentNode.Children = append(parentNode.Children, nodes[child])
		}
	}

	return nodes
}

func dfsQuery(tree map[int]*DirectoryNode, n int) int64 {
	if tree == nil {
		return -1
	}

	var dfs func(node *DirectoryNode) int64
	dfs = func(node *DirectoryNode) int64 {
		result := node.Size
		for _, child := range node.Children {
			result += dfs(child)
		}
		return result
	}
	return dfs(tree[n])
}

func calculater(n int, arr []string) int64 {
	if len(arr) == 0 {
		return 0
	}

	tree := buildTreeMap(arr)
	return dfsQuery(tree, n)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		strArr := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		m, _ := strconv.Atoi(strArr[0])
		n, _ := strconv.Atoi(strArr[1])

		arr := make([]string, 0, m)
		count := 0
		for count < m {
			if !scanner.Scan() {
				break
			}
			arr = append(arr, scanner.Text())
			count++
		}
		total := calculater(n, arr)
		fmt.Println(total)
	}
}
