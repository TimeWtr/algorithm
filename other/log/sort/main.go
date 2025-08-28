// 题目描述
// [运维工程师]采集到某产品线网运行一天产生的日志n条，现需根据日志时间先后顺序对日志进行排序，日志时间格式为H:M:S.N。

// H表示小时(0~23)
// M表示分钟(0~59)
// S表示秒(0~59)
// N表示毫秒(0~999)
// 时间可能并没有补全，也就是说，01:01:01.001也可能表示为1:1:1.1。

// 输入描述
// 第一行输入一个整数n表示日志条数，1<=n<=100000，接下来n行输入n个时间。

// 输出描述
// 按时间升序排序之后的时间，如果有两个时间表示的时间相同，则保持输入顺序。

// 示例1
// 输入

// 2
// 01:41:8.9
// 1:1:09.211
// 1
// 2
// 3
// 输出

// 1:1:09.211
// 01:41:8.9
// 1
// 2
// 说明

// 示例2
// 输入

// 3
// 23:41:08.023
// 1:1:09.211
// 08:01:22.0
// 1
// 2
// 3
// 4
// 输出

// 1:1:09.211
// 08:01:22.0
// 23:41:08.023
// 1
// 2
// 3
// 说明

// 示例3
// 输入

// 2
// 22:41:08.023
// 22:41:08.23
// 1
// 2
// 3
// 输出

// 22:41:08.023
// 22:41:08.23
//
// 解题思路：
// 将日志时间进行分割，小时、分钟、秒、毫秒都转换成毫秒，计算总毫秒数，根据总毫秒数来进行排序
// 总毫秒数 = 小时*3600*1000 + 分钟*60*1000 + 秒*1000 + 毫秒

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 日志条目
type LogEntry struct {
	// 原始时间字符串
	TimeStr string
	// 毫秒时间戳
	Ts int64
}

func Sort(arr []string, n int) []LogEntry {
	if len(arr) == 0 {
		return nil
	}

	res := make([]LogEntry, 0, n)
	for _, str := range arr {
		strArr := strings.Split(str, ".")
		timePart := strArr[0]
		msPart := "0"
		if len(strArr) > 1 {
			msPart = strArr[1]
		}

		timeUnits := strings.Split(timePart, ":")
		hours, _ := strconv.Atoi(timeUnits[0])
		minutes, _ := strconv.Atoi(timeUnits[1])
		seconds, _ := strconv.Atoi(timeUnits[2])
		mills, _ := strconv.Atoi(msPart)
		totalMills := hours*3600*1000 + minutes*60*1000 + seconds*1000 + mills
		res = append(res, LogEntry{
			TimeStr: str,
			Ts:      int64(totalMills),
		})
	}

	// 切片排序
	sort.SliceStable(res, func(i, j int) bool {
		return res[i].Ts < res[j].Ts
	})

	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		arr := make([]string, 0, n)
		count := 0
		for count < n {
			count++
			if !scanner.Scan() {
				break
			}
			t := strings.TrimSpace(scanner.Text())
			if t == "" {
				break
			}

			arr = append(arr, t)
		}
		sortedArr := Sort(arr, n)
		for _, val := range sortedArr {
			fmt.Println(val.TimeStr)
		}
	}
}
