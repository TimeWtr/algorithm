// 题目：数字分类
// 题目描述
// 对一个数据a进行分类，分类方法为：

// 此数据a（四个字节大小）的四个字节相加对一个给定的值b取模，如果得到的结果小于一个给定的值c，
//  则数据a为有效类型，其类型为取模的值；如果得到的结果大于或者等于c，则数据a为无效类型。

// 比如一个数据a=0x01010101，b=3，按照分类方法计算（0x01+0x01+0x01+0x01）%3=1，

// 所以如果c=2，则此a为有效类型，其类型为1，如果c=1，则此a为无效类型；

// 又比如一个数据a=0x01010103，b=3，按照分类方法计算（0x01+0x01+0x01+0x03）%3=0，

// 所以如果c=2，则此a为有效类型，其类型为0，如果c=0，则此a为无效类型。

// 输入12个数据，第一个数据为c，第二个数据为b，剩余10个数据为需要分类的数据，

// 请找到有效类型中包含数据最多的类型，并输出该类型含有多少个数据。

// 输入描述
// 输入12个数据，用空格分隔，第一个数据为c，第二个数据为b，剩余10个数据为需要分类的数据。

// 输出描述
// 输出最多数据的有效类型有多少个数据。

// 用例1
// 输入
// 3 4 256 257 258 259 260 261 262 263 264 265
// 输出
// 3
// 说明
// 10个数据4个字节相加后的结果分别为1 2 3 4 5 6 7 8 9 10，

// 故对4取模的结果为1 2 3 0 1 2 3 0 1 2，c为3，所以0 1 2都是有效类型，类型为1和2的有3个数据，类型为0的只有2个数据，故输出3。

// 用例2
// 输入
// 1 4 256 257 258 259 260 261 262 263 264 265
// 输出
// 2
// 说明
// 10个数据4个字节相加后的结果分别为1 2 3 4 5 6 7 8 9 10，

// 故对4取模的结果为1 2 3 0 1 2 3 0 1 2，c为1，

// 所以只有0是有效类型，类型为0的有2个数据，故输出2。

// 题解
// 思路：一个简单的位运算题目。

// 题目的意思是将四字节分割成四个一字节的16进制，然后进行累加，取余数。
// 使用位右移 8位进行快速进行分割字节，累加。
// 本题额外注意，32位对于int可能会超范围。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handler(c, b int, arr []int) int {
	if len(arr) != 10 {
		return -1
	}

	validArr := []int32{}
	for _, v := range arr {
		sum := getSum(int32(v))
		res := sum % int32(b)
		if res < int32(c) {
			validArr = append(validArr, res)
		}
	}

	// 分组统计
	group := map[int32]int{}
	for _, v := range validArr {
		group[v]++
	}

	maxCount := -1
	for _, v := range group {
		if v > maxCount {
			maxCount = v
		}
	}

	return maxCount
}

// getSum 通过位操作计算4字节数字的和
func getSum(val int32) int32 {
	byte1 := val & 0xFFF
	byte2 := (val >> 8) & 0xFFF
	byte3 := (val >> 16) & 0xFFF
	byte4 := (val >> 24) & 0xFFF
	return byte1 + byte2 + byte3 + byte4
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strArr := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		if len(strArr) != 12 {
			fmt.Println("数据条数不正确")
			break
		}

		c, _ := strconv.Atoi(strArr[0])
		b, _ := strconv.Atoi(strArr[1])
		arr := make([]int, 0, 10)
		for _, v := range strArr[2:] {
			val, _ := strconv.Atoi(v)
			arr = append(arr, val)
		}
		res := handler(c, b, arr)
		fmt.Println(res)
	}
}
