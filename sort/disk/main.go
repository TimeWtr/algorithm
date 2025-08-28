// 题目：磁盘容量排序
// 题目描述
// 磁盘的容量单位常用的有 M，G，T 这三个等级，它们之间的换算关系为：

// 1T = 1024G
// 1G = 1024M
// 现在给定 n 块磁盘的容量，请对它们按从小到大的顺序进行稳定排序。

// 例如给定5块盘的容量：

// 1T，20M，3G，10G6T，3M12G9M

// 排序后的结果为：

// 20M，3G，3M12G9M，1T，10G6T

// 注意单位可以重复出现，上述 3M12G9M 表示的容量即为：3M+12G+9M，和 12M12G 相等。

// 输入描述
// 输入第一行包含一个整数 n，表示磁盘的个数

// 2 ≤ n ≤ 100
// 接下的 n 行，每行一个字符串（长度大于2，小于30），表示磁盘的容量，由一个或多个格式为mv的子串组成，其中 m 表示容量大小，v 表示容量单位，例如：20M，1T，30G，10G6T，3M12G9M。

// 磁盘容量 m 的范围为 1 到 1024 的正整数
// 容量单位 v 的范围只包含题目中提到的 M，G，T 三种，换算关系如题目描述
// 输出描述
// 输出 n 行，表示 n 块磁盘容量排序后的结果。

// 示例1
// 输入

// 3
// 1G
// 2G
// 1024M
// 1
// 2
// 3
// 4
// 输出

// 1G
// 1024M
// 2G
// 1
// 2
// 3
// 说明

// 1G和1024M容量相等，稳定排序要求保留它们原来的相对位置，故1G在1024M之前。

// 示例2
// 输入

// 3
// 2G4M
// 3M2G
// 1T
// 1
// 2
// 3
// 4
// 输出

// 3M2G
// 2G4M
// 1T
// 1
// 2
// 3
// 说明

// 1T的容量大于2G4M，2G4M的容量大于3M2G。

package main

import (
	"bufio"
	"fmt"
	"os"
	st "sort"
	"strconv"
	"strings"
)

type DiskEntry struct {
	// 原始的字符串
	str string
	// 转换后的int64容量
	cp int64
}

func sort(arr []string) []string {
	if len(arr) == 0 {
		return nil
	}

	disks := make([]DiskEntry, len(arr))
	// 解析字符串获取到具体的容量
	for i, str := range arr {
		disks[i] = DiskEntry{
			str: str,
			cp:  parseStr(str),
		}
	}

	// 排序
	st.SliceStable(disks, func(i int, j int) bool {
		return disks[i].cp < disks[j].cp
	})

	// 构建结果
	res := make([]string, len(disks))
	for _, v := range disks {
		res = append(res, v.str)
	}

	return res
}

// parseStr 解析磁盘空间字符串，并转换为MB单位
func parseStr(str string) int64 {
	total := 0
	numStr := ""
	for _, v := range str {
		if v >= '0' && v <= '9' {
			numStr += string(v)
		} else {
			n, _ := strconv.Atoi(numStr)
			switch string(v) {
			case "M":
				// 当前容量单位为MB
				total += n
			case "G":
				// 当前单位为GB，转换为MB
				total += 1024 * n
			case "T":
				total += 1024 * 1024
			}
			numStr = ""
		}
	}

	return int64(total)
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
			if !scanner.Scan() {
				break
			}
			capStr := strings.TrimSpace(scanner.Text())
			arr = append(arr, capStr)
			count++
		}
		res := sort(arr)
		for _, v := range res {
			fmt.Println(v)
		}
	}
}
