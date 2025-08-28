// 题目描述
// 现有两个整数数组，需要你找出两个数组中同时出现的整数，并按照如下要求输出:有同时出现的整数时，先按照同时出现次数（整数在两个数组中都出现并目出现次数较少的那个）进行归类，然后按照出现次数从小到大依次按行输出。 没有同时出现的整数时，输出NULL.

// 输入描述
// 第一行为第一个整数数组，第二行为第二个整数数组，每行数中整数与整数之间以英文号分，整数的取值范用为[-200, 200]，数组长度的范用为[1, 10000]之间的整数.

// 输出描述
// 按照出现次数从小到大依次按行输出，每行输出的格式为：

// 出现次数:该出现次数下的整数升序排序的结果

// 格式中的":"为英文冒号，整数间以英文逗号分隔

// 用例1
// 输入
// 5,3,6,-8,0,11
// 2,8,8,8,-1,15
// 输出
// NULL
// 用例2
// 输入
// 5,8,11,3,6,8,8,-1,11,2,11,11
// 11,2,11,8,6,8,8,-1,8,15,3,-9,11
// 输出
// 1:-1,2,3,6
// 3:8,11
// 说明
// 说明 两整数数组中同时出现的整数为-12、3、6、8、11,其中同时出现次数为1的整数为-1,2,3,6(升序排序),同时出现次数为3的整数为8,11(升序排序),先升序输出出现次数为1的整数，再升序输出出现次数为3的整数

// 解题思路：
// 1. 创建两个map，分别统计两个数组中的整数的重复次数
// 2. 遍历两个map统计两个数组中同时出现的整数的次数(比较并获取较小的出现次数)
// 3. 分组：根据出现的次数进行分组统计，即出现一次的整数，出现两次的整数...
// 4. 排序：每一个次数对应一个整数数组，对数组进行升序排序
package main

import (
	"bufio"
	"fmt"
	"os"
	st "sort"
	"strconv"
	"strings"
)

// getIntArr 转换int数组
func getIntArr(strArr []string) []int {
	if len(strArr) == 0 {
		return nil
	}

	arr := make([]int, len(strArr))
	for i, v := range strArr {
		val, _ := strconv.Atoi(v)
		arr[i] = val
	}

	return arr
}

func sort(arr1, arr2 []int) []string {
	if len(arr1) == 0 || len(arr2) == 0 {
		return nil
	}

	// 统计两个数组中的重复整数次数
	// key: 整数，val：出现的次数
	m1 := map[int]int{}
	m2 := map[int]int{}
	for _, v := range arr1 {
		m1[v]++
	}

	for _, v := range arr2 {
		m2[v]++
	}

	// 分组：找出两个数组中的重复整数和次数
	// key: 重复的次数，val：重复的整数数组
	repeatSet := map[int][]int{}
	for k1, v1 := range m1 {
		v2, ok := m2[k1]
		if !ok {
			continue
		}

		// 重复的整数
		minCount := min(v1, v2)
		set, ok := repeatSet[minCount]
		if !ok {
			set = []int{}
		}
		set = append(set, k1)
		st.Ints(set)
		repeatSet[minCount] = set
	}

	if len(repeatSet) == 0 {
		return nil
	}

	// 使用单独的数组存储重复的次数，并进行排序
	sortedKeySet := make([]int, 0, len(repeatSet))
	for k := range repeatSet {
		sortedKeySet = append(sortedKeySet, k)
	}
	st.Ints(sortedKeySet)

	// 按照排序后的顺序生成输出的结果
	res := make([]string, 0, len(sortedKeySet))
	for _, k := range sortedKeySet {
		valSet := repeatSet[k]
		str := strconv.Itoa(k) + ":"
		for i, v := range valSet {
			str += strconv.Itoa(v)
			if i < len(valSet)-1 {
				str += ","
			}
		}
		res = append(res, str)
	}

	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		firstStrArr := strings.Split(strings.TrimSpace(scanner.Text()), ",")
		firstArr := getIntArr(firstStrArr)

		if !scanner.Scan() {
			break
		}
		secondStrArr := strings.Split(strings.TrimSpace(scanner.Text()), ",")
		secondArr := getIntArr(secondStrArr)
		res := sort(firstArr, secondArr)
		if res == nil || len(res) == 0 {
			fmt.Println("NULL")
		} else {
			for _, v := range res {
				fmt.Println(v)
			}
		}
	}
}
