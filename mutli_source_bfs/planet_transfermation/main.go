// 题目：星球改造计划
// 题目描述
//2XXX年，人类通过对火星的大气进行宜居改造分析，使得火星已在理论上具备人类宜居的条件；
//
//由于技术原因，无法一次性将火星大气全部改造，只能通过局部处理形式；
//
//假设将火星待改造的区域为row *
//column的网格，每个网格有3个值，宜居区、可改造区、死亡区，使用YES、NO、NA代替，YES表示该网格已经完成大气改造，NO表示该网格未进行改造，后期可进行改造，NA表示死亡区，不作为判断是否改造完的宜居，无法穿过；
//
//初始化下，该区域可能存在多个宜居区，并目每个宜居区能同时在每个大阳日单位向上下左右四个方向的相邻格子进行扩散，自动将4个方向相邻的真空区改造成宜居区；
//
//请计算这个待改造区域的网格中，可改造区是否能全部成宜居区，如果可以，则返回改造的大阳日天教，不可以则返回-1
//
//输入描述
//输入row * column个网格数据，每个网格值枚举值如下: YES，NO，NA；
//
//样例:、
//
//YES YES NO
//NO NO NO
//NA NO YES
//1
//2
//3
//备注
//grid[i][j]只有3种情况，YES、NO、NA
//
//row == grid.length
//column == grid[i].length
//1 ≤ row, column ≤ 8
//输出描述
//可改造区是否能全部变成宜居区，如果可以，则返回改造的太阳日天数，不可以则返回-1。
//
//示例1
//输入
//
//YES YES NO
//NO NO NO
//YES NO NO
//1
//2
//3
//输出
//
//2
//1
//说明
//
//经过 2 个太阳日，完成宜居改造。
//
//示例2
//输入
//
//YES NO NO NO
//NO NO NO NO
//NO NO NO NO
//NO NO NO NO
//1
//2
//3
//4
//输出
//
//6
//1
//说明
//
//经过 6 个太阳日，可完成改造
//
//示例3
//输入
//
//NO NA
//1
//输出
//
//-1
//1
//说明
//
//无改造初始条件，无法进行改造
//
//示例4
//输入
//
//YES NO NO YES
//NO NO YES NO
//NO YES NA NA
//YES NO NA NO
//1
//2
//3
//4
//输出
//
//-1
//1
//说明
//
//-1 ，右下角的区域，被周边三个死亡区挡住，无法实现改造

// 解法：构建二维矩阵，标记所有的YES为0，NO为可改造点，初始状态为-1
// 标识为可改造且还未改造，使用多源BFS从所有为0的源点开始BFS遍历和标记。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getGrid(rows []string) [][]string {
	grid := make([][]string, len(rows))
	for i, row := range rows {
		fields := strings.Fields(row)
		grid[i] = fields
	}

	return grid
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}

		num, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		strGrid := make([]string, 0, num)
		count := 0
		for count < num {
			for !scanner.Scan() {
				break
			}
			strGrid = append(strGrid, strings.TrimSpace(scanner.Text()))
			count++
		}

		// 构建网格
		grid := getGrid(strGrid)

		// 初始化距离矩阵，-1标识未被访问过
		dist := make([][]int, len(grid))
		for i := range dist {
			dist[i] = make([]int, len(grid[i]))
			for j := range dist[i] {
				dist[i][j] = -1
			}
		}

		// 定义方向坐标，计算机坐标
		direction := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		var queue [][2]int
		// 将所有为YES的位置的距离设置为0，并将该坐标坐标加入到队列中
		for i, col := range grid {
			for j := range col {
				if col[j] == "YES" {
					dist[i][j] = 0
					queue = append(queue, [2]int{i, j})
				}
			}
		}

		// 处理没有YES的情况
		if len(queue) == 0 {
			for i := range grid {
				for j := range grid[i] {
					if grid[i][j] == "NO" {
						// 没有宜居区且还存在未开发区，这些待开发区将永远无法开发
						fmt.Println(-1)
						return
					}
				}
			}

			// 没有宜居区且不存在待开发区，所以开发期为0
			fmt.Println(0)
			return
		}

		// 多源BFS进行扩散计算
		maxDays := 0
		for len(queue) > 0 {
			cell := queue[0]
			queue = queue[1:]
			i, j := cell[0], cell[1]
			for _, dic := range direction {
				ni, nj := i+dic[0], j+dic[1]
				// 坐标在矩阵范围内
				if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[0]) {
					if grid[ni][nj] == "NO" && dist[ni][nj] == -1 {
						// 邻居点的改造时间是当前点的基础上+1，这是标准的扩散逻辑，
						// 外层是内层的基础上计算，每多一层就改造天数+1
						dist[ni][nj] = dist[i][j] + 1
						if dist[ni][nj] > maxDays {
							maxDays = dist[ni][nj]
						}
						// 将新的改造点加入队列中
						queue = append(queue, [2]int{ni, nj})
					}
				}
			}
		}

		// 检查是否有待改造区没有被改造
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j] == "NO" && dist[i][j] == -1 {
					fmt.Println(-1)
					return
				}
			}
		}
		fmt.Println(maxDays)
	}
}
