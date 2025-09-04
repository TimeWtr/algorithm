// 解题思路：直接构建树会造成内存占用和时间占用较多，可以使用隐式的实现树的关系
// 需要两个哈希表来存储目录信息和映射关系。
// 哈希表1: 存储节点ID和节点大小的映射关系，key为节点ID，val为目录节点的大小
// 哈希表2: 存储目录节点与子节点的映射关系，key为节点ID， val为子节点的ID组成的数组
// 使用DFS算法递归计算并返回指定节点目录的总大小

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculater(arr []string, n int) int64 {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	// 创建节点大小映射关系
	nodes := make(map[int]int64)
	// 构建节点和子节点的映射关系
	nodesRelationship := make(map[int][]int)
	for _, line := range arr {
		fields := strings.Fields(line)
		id, _ := strconv.Atoi(fields[0])
		size, _ := strconv.ParseInt(fields[1], 10, 64)
		nodes[id] = size

		// 处理子节点的映射关系
		part := fields[2]
		if part != "()" {
			part = strings.TrimSuffix(strings.TrimPrefix(part, "("), ")")
			children := strings.Split(part, ",")
			for _, child := range children {
				intChild, _ := strconv.Atoi(child)
				nodesRelationship[id] = append(nodesRelationship[id], intChild)
			}
		}
	}

	// 计算目录总大小
	var dfs func(dirID int) int64
	dfs = func(dirID int) int64 {
		result := nodes[dirID]
		for _, child := range nodesRelationship[dirID] {
			result += dfs(child)
		}
		return result
	}
	return dfs(n)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		fields := strings.Fields(scanner.Text())
		m, _ := strconv.Atoi(fields[0])
		n, _ := strconv.Atoi(fields[1])

		arr := make([]string, 0, m)
		count := 0
		for count < m {
			if !scanner.Scan() {
				break
			}
			arr = append(arr, strings.TrimSpace(scanner.Text()))
			count++
		}
		res := calculater(arr, n)
		fmt.Println(res)
	}
}
