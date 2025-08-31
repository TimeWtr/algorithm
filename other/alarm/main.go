// 题目：告警抑制
// 题目描述
// 告警抑制，是指高优先级告警抑制低优先级告警的规则。高优先级告警产生后，低优先级告警不再产生。请根据原始告警列表和告警抑制关系，给出实际产生的告警列表。

// 不会出现循环抑制的情况。
// 告警不会传递，比如A->B,B->C，这种情况下A不会直接抑制C。但被抑制的告警仍然可以抑制其他低优先级告警。
// 输入描述
// 第一行为数字N，表示告警抑制关系个数，0 ≤ N ≤ 120
// 接下来N行，每行是由空格分隔的两个告警ID，例如: id1 id2，表示id1抑制id2，告警ID的格式为：

// 大写字母+0个或者1个数字

// 最后一行为告警产生列表，列表长度[1,100]

// 输出描述
// 真实产生的告警列表

// 备注
// 告警ID之间以单个空格分隔

// 用例1
// 输入
// 2
// A B
// B C
// A B C D E
// 输出
// A D E
// 说明
// A抑制了B，B抑制了C，最后实际的告警为A D E

// 用例2
// 输入
// 4
// F G
// C B
// A G
// A0 A
// A B C D E
// 输出
// A C D E

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handler(conditions, arr []string) []string {
	// 边界条件
	if len(arr) == 0 || len(conditions) == 0 {
		return arr
	}

	// 构建抑制关系
	edges := make([][]string, 0, len(conditions))
	for _, v := range conditions {
		conArr := strings.Split(v, " ")
		edges = append(edges, []string{conArr[0], conArr[1]})
	}

	// 构建告警集合
	alarmSet := make(map[string]bool)
	for _, v := range arr {
		alarmSet[v] = true
	}

	// 构建告警抑制有向图
	// key: 告警ID，value：告警ID对应的被抑制告警ID列表，一个告警ID可能对应多个被抑制告警ID
	graph := make(map[string][]string)
	for _, v := range edges {
		alarm, alarmed := v[0], v[1]
		// 只有当告警ID和被告警ID都在告警列表才可以抑制生效
		if alarmSet[alarm] && alarmSet[alarmed] {
			set, ok := graph[alarm]
			if !ok {
				set = []string{}
			}
			set = append(set, alarmed)
			graph[alarm] = set
		}
	}

	// 构建被抑制的告警集合
	inhibitedSet := make(map[string]bool)
	for _, start := range arr {
		if _, exist := graph[start]; !exist {
			continue
		}

		// BFS初始化
		visited := make(map[string]bool)
		queue := []string{start}
		visited[start] = true
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]

			// 当前节点不是起始节点，需要标记为被抑制节点
			if curr != start {
				inhibitedSet[curr] = true
			}

			// 遍历当前节点的所有邻居
			for _, next := range graph[curr] {
				if !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
	}

	// 构建未被抑制的结果
	res := []string{}
	for _, v := range arr {
		if !inhibitedSet[v] {
			res = append(res, v)
		}
	}

	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		conditions := make([]string, 0, n)
		count := 0
		for count < n {
			if !scanner.Scan() {
				break
			}
			conditions = append(conditions, strings.TrimSpace(scanner.Text()))
			count++
		}

		arr := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		res := handler(conditions, arr)
		fmt.Println(res)
	}
}
