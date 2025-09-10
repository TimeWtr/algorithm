// 题目：取最少量的球
// 题目描述
//某部门开展 Family Day 开放日活动，其中有个从桶里取球的游戏，游戏规则如下：
// 有 N 个容量一样的小桶等距排开，且每个小桶都默认装了数量不等的小球，每个
// 小桶装的小球数量记录在数组 bucketBallNums 中，游戏开始时，要求所有桶的小
// 球总数不能超过 SUM，如果小球总数超过 SUM，则需对所有的小桶统一设置一个容
// 量最大值 maxCapacity，并需将超过容量最大值的小球拿出来，直至小桶里的小球数
// 量小于 maxCapacity。请您根据输入的数据，计算从每个小桶里拿出的小球数量？
// 限制规则一：所有小桶的小球总和小于 SUM，则无需设置容量值 maxCapacity，并且
// 无需从小桶中拿球出来，返回结果[] 限制规则二：如果所有小桶的小球总和大于 SUM，
// 则需设置容量最大值 maxCapacity，并且需从小桶中拿球出来，返回从每个小桶拿出的
// 小球数量组成的数组
//
//输入描述
//第一行输入 2 个正整数，数字之间使用空格隔开，其中：第一个数字表示 SUM 第二个数字表示 bucketBallNums 数组长度.
//第二行输入 N 个正整数，数字之间使用空格隔开，表示 bucketBallNums 的每一项
//备注
//1 ≤ bucketBallNums[i] ≤ 10^9 1 ≤ bucketBallNums.length = N ≤ 10^5 1 ≤ maxCapacity ≤ 10^9 1 ≤ SUM ≤ 10^9
//
//输出描述
//从每个小桶里拿出的小球数量，并使用一维数组表示
//用例1
//输入
//14 7
//2 3 2 5 5 1 4
//输出
//[0,1,0,3,3,0,2]
//用例2
//输入
//3 3
//1 2 3
//输出
//[0,1,2]
//用例3
//输入
//6 2
//3 2
//输出
//[]

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
		fields := strings.Fields(strings.TrimSpace(scanner.Text()))
		sum, _ := strconv.Atoi(fields[0])
		num, _ := strconv.Atoi(fields[1])

		if !scanner.Scan() {
			break
		}

		buckets := make([]int, num)
		bucketFields := strings.Fields(strings.TrimSpace(scanner.Text()))
		for i := range bucketFields {
			buckets[i], _ = strconv.Atoi(bucketFields[i])
		}

		// 统计所有的桶中的数量和是否超过最大和限制，如果不超过则直接打印
		// 如果超过则计算出所有桶中的数量最大值
		maxVal := 0
		total := 0
		for i := range buckets {
			if buckets[i] > maxVal {
				maxVal = buckets[i]
			}
			total += buckets[i]
		}

		if total <= sum {
			fmt.Println("[]")
			break
		}

		// 二分查找计算出最大的容量
		left, right := 0, maxVal
		ansCap := 0
		for left <= right {
			mid := left + (right-left)/2
			sumTemp := 0
			for _, bk := range buckets {
				if bk > mid {
					sumTemp += mid
				} else {
					sumTemp += bk
				}
			}

			if sumTemp <= sum {
				ansCap = mid
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		// 根据找到的最大容量从每一个超过的桶中取小球
		results := make([]int, num)
		for i, bk := range buckets {
			if bk > ansCap {
				results[i] = bk - ansCap
			} else {
				results[i] = 0
			}
		}

		// 格式化输出结果
		var builder strings.Builder
		builder.WriteString("[")
		for i := range results {
			builder.WriteString(strconv.Itoa(results[i]))
			if i != len(results)-1 {
				builder.WriteString(", ")
			}
		}
		builder.WriteString("]")
		fmt.Println(builder.String())
	}
}
