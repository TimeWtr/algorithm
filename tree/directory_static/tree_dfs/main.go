// 解题思路：这是一个树型结构，子节点在()内，可以先存储每一个目录节点的属性
// 再创建当前节点与子节点的映射关系，根据节点属性和映射关系显式的构建一颗树，
// 找到对应的节点，然后根据DFS算法计算目录的总大小。
// 优点：显式的构建树，直观易于理解
// 缺点：构建树需要占用较大的内存空间和时间

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
