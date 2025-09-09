// 题目：光伏场地建设规划
//题目描述
//祖国西北部有一片大片荒地，其中零星的分布着一些湖泊，保护区，矿区;
//整体上常年光照良好，但是也有一些地区光照不太好。
//
//某电力公司希望在这里建设多个光伏电站，生产清洁能源对每平方公里的土地进行了发电评估，
//其中不能建设的区域发电量为0kw，可以发电的区域根据光照，地形等给出了每平方公里年发电量x千瓦。
//我们希望能够找到其中集中的矩形区域建设电站，能够获得良好的收益。
//
//输入描述
//第一行输入为调研的地区长，宽，以及准备建设的电站【长宽相等，为正方形】的边长最低要求的发电量
//之后每行为调研区域每平方公里的发电量
//
//输出描述
//输出为这样的区域有多少个
//
//示例1
//输入
//
//2 5 2 6
//1 3 4 5 8
//2 3 6 7 1
//1
//2
//3
//输出
//
//4
//1
//说明
//
//输入含义：
//调研的区域大小为长2宽5的矩形，我们要建设的电站的边长为2，建设电站最低发电量为6.
//输出含义：
//长宽为2的正方形满足发电量大于等于6的区域有4个。
//
//示例2
//输入
//
//5 1 6
//1 3 4 5 8
//2 3 6 7 1
//1
//2
//3
//输出
//
//3
//1
//说明
//
//解题思路
//本题可以使用动态规划前缀和思想解题。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())
		fields := strings.Fields(line)
		m, _ := strconv.Atoi(fields[0])
		n, _ := strconv.Atoi(fields[1])
		c, _ := strconv.Atoi(fields[2])
		k, _ := strconv.Atoi(fields[3])
		if c > m || c > n {
			// 需要的光伏场地变长大于实际的场地大小
			fmt.Println(0)
			break
		}

		// 读取发电量的二维矩阵
		grid := make([][]int, m)
		for i := 0; i < m; i++ {
			if !scanner.Scan() {
				break
			}
			grid[i] = make([]int, n)
			strFields := strings.Fields(strings.TrimSpace(scanner.Text()))
			for j := 0; j < len(strFields); j++ {
				grid[i][j], _ = strconv.Atoi(strFields[j])
			}
		}

		// 使用滑动窗口计算二维矩阵每一行的横向滑动窗口和
		rowSums := make([][]int, m)
		for i := 0; i < m; i++ {
			// 定义第1行的和矩阵，窗口数量为n-c+1，即元素数量-窗口大小+1
			rowSums[i] = make([]int, n-c+1)
			sum := 0
			// 计算第一个窗口
			for j := 0; j < c; j++ {
				sum += grid[i][j]
			}
			rowSums[i][0] = sum

			// 基于第一个窗口继续向后滑动计算所有的窗口
			// j代表的是窗口开始的元素下标，每次滑动在之前的sum基础上
			// 减去前一个过期的元素，再加上后一个刚进入窗口的元素
			for j := 1; j <= n-c; j++ {
				sum = sum - grid[i][j-1] + grid[i][j+c-1]
				rowSums[i][j] = sum
			}
		}

		// 现在对二维矩阵和进行竖向的滑动窗口统计
		count := 0
		for j := 0; j < n-c+1; j++ {
			sum := 0
			// 计算第一个滑动窗口的大小
			for i := 0; i < c; i++ {
				sum += rowSums[i][j]
			}

			if sum >= k {
				// 电力满足要求
				count++
			}

			for i := 1; i <= m-c; i++ {
				sum = sum - rowSums[i-1][j] + rowSums[i][j]
				if sum >= k {
					count++
				}
			}
		}

		fmt.Println(count)
	}
}
