// 题目：攀登者
// 题目描述
//攀登者喜欢寻找各种地图，并且尝试攀登到最高的山峰。
//
//地图表示为一维数组，数组的索引代表水平位置，数组的元素代表相对海拔高度。其中数组元素0代表地面。
//
//例如：[0,1,2,4,3,1,0,0,1,2,3,1,2,1,0]，代表如下图所示的地图，地图中有两个山脉位置分别为 1,2,3,4,5 和 8,9,10,11,12,13，最高峰高度分别为 4,3。最高峰位置分别为3,10。
//
//一个山脉可能有多座山峰(高度大于相邻位置的高度，或在地图边界且高度大于相邻的高度)。
//
//63287f530c7641628550cf45ab259f28.png
//
//登山时会消耗登山者的体力(整数)，
//
//上山时，消耗相邻高度差两倍的体力
//下山时，消耗相邻高度差一倍的体力
//平地不消耗体力
//登山者体力消耗到零时会有生命危险。
//
//例如，上图所示的山峰：
//
//从索引0，走到索引1，高度差为1，需要消耗 2 * 1 = 2 的体力，
//从索引2，走到索引3，高度差为2，需要消耗 2 * 2 = 4 的体力。
//从索引3，走到索引4，高度差为1，需要消耗 1 * 1 = 1 的体力。
//
//
//攀登者想要评估一张地图内有多少座山峰可以进行攀登，且可以安全返回到地面，且无生命危险。
//
//例如上图中的数组，有3个不同的山峰，登上位置在3的山可以从位置0或者位置6开始，从位置0登到山顶需要消耗体力 1 * 2 + 1 * 2 + 2 * 2 = 8，从山顶返回到地面0需要消耗体力 2 * 1 + 1 * 1 + 1 * 1 = 4 的体力，按照登山路线 0 → 3 → 0 需要消耗体力12。攀登者至少需要12以上的体力（大于12）才能安全返回。
//
//输入描述
//第一行输入为地图一维数组
//
//第二行输入为攀登者的体力
//
//输出描述
//确保可以安全返回地面，且无生命危险的情况下，地图中有多少山峰可以攀登。
//
//用例1
//输入
//0,1,4,3,1,0,0,1,2,3,1,2,1,0
//13
//输出
//3
//说明
//登山者只能登上位置10和12的山峰，7 → 10 → 7，14 → 12 → 14
//
//用例2
//输入
//1,4,3
//999
//输出
//0
//说明
//没有合适的起点和终点
//
//

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// getGrounds 获取所有平地的下标
func getGrounds(heights []int) []int {
	var grounds []int
	for index, height := range heights {
		if height == 0 {
			grounds = append(grounds, index)
		}
	}

	return grounds
}

// getPeeks 遍历整个数组，标记出山顶和非山顶区域，山顶部分为true
// 非山顶区域标记为false，山顶区域的条件是：
// 1. 如果在边界区域，则左边界的高度应该比后续的高度高，右边界的高度应该比前边的高度高
// 2. 如果是中间区域，则需要满足，当前的高度均高于前后两个相邻的区域的高度
func getPeeks(heights []int) []bool {
	peeks := make([]bool, len(heights))
	n := len(heights)
	for i := range heights {
		leftOk := true
		if i > 0 {
			// 中间区域，判断左半部分
			if heights[i-1] >= heights[i] {
				leftOk = false
			}
		}

		rightOk := true
		if i < n-1 {
			// 中间区域，判断右半部分
			if heights[i+1] >= heights[i] {
				rightOk = false
			}
		}

		if leftOk && rightOk {
			peeks[i] = true
		}
	}

	return peeks
}

// costFromAtoB A和B都是下标，计算从A到B耗费的体力，注意这里边的A和B都可以是山顶或者平地
// 比如：A是平地时B就是山顶，A是山顶时B就是平地，对应的路线就是上山和下山
// 遍历中间的高度计算消耗的体力值
func costFromAtoB(a, b int, heights []int) int {
	cost := 0
	if a < b {
		// 从左往右
		for i := a; i < b; i++ {
			diff := heights[i+1] - heights[i]
			if diff > 0 {
				// 上山路线
				cost += diff * 2
			} else {
				// 下山路线
				cost += -diff * 1
			}
		}
	} else {
		// 从右往左
		for i := b; i > a; i-- {
			diff := heights[i-1] - heights[i]
			if diff > 0 {
				// 上山路线
				cost += diff * 2
			} else {
				// 下山路线
				cost += -diff * 1
			}
		}
	}

	return cost
}

func transferIntArr(strArr []string) []int {
	arr := make([]int, len(strArr))
	for i := range strArr {
		arr[i], _ = strconv.Atoi(strArr[i])
	}

	return arr
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		heightsStrArr := strings.Split(strings.TrimSpace(scanner.Text()), ",")
		heightsArr := transferIntArr(heightsStrArr)

		if !scanner.Scan() {
			break
		}
		strength, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		// 找到所有的高度为0的下标(平地)
		grounds := getGrounds(heightsArr)
		// 标记山顶和平地
		peeks := getPeeks(heightsArr)
		// 遍历计算耗费的体力
		count := 0
		for i := range peeks {
			if !peeks[i] {
				// 不是山顶区域
				continue
			}

			// 最小的体力
			minStrength := 1 << 30
			// 找出所有的平地
			for _, g := range grounds {
				// 上山耗费的体力
				costA := costFromAtoB(g, i, heightsArr)
				// 下山耗费的体力
				costB := costFromAtoB(i, g, heightsArr)
				total := costA + costB
				if total < minStrength {
					minStrength = total
				}
			}

			if minStrength < strength {
				count++
			}
		}
		fmt.Println(count)
	}
}
