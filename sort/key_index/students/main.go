// 学生分组排序
// 题目描述：
// 某班级有 n 个学生，每个学生有一个唯一学号和一个小组编号（1-5）。
// 请按照小组编号从小到大对学生进行排序，要求相同小组的学生保持原来的相对顺序（稳定排序）。
//
// 输入格式：
// 第一行：整数 n，表示学生数量

// 接下来 n 行：每行包含一个学号和一个小组编号，用空格分隔
// 输出格式：

// 排序后的学生信息，每行一个学生的学号和小组编号
// 示例输入：
// 8
// 1001 3
// 1002 2
// 1003 4
// 1004 1
// 1005 2
// 1006 3
// 1007 5
// 1008 1
//
// 示例输出：
// 1004 1
// 1008 1
// 1002 2
// 1005 2
// 1001 3
// 1006 3
// 1003 4
// 1007 5
//
// 解题思路：使用键索引计数排序算法

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	// 学号
	ID int
	// 分组
	Group int
}

func keyIndexSort(students []Student) []string {
	if len(students) == 0 {
		return nil
	}

	l := len(students)
	// 统计每一个小组的学生数量，容量多2个，方便计算，索引必须+1
	counts := make([]int, l+2)
	for _, student := range students {
		counts[student.Group+1]++
	}

	// 将频率转换为起始索引
	for i := 0; i <= l; i++ {
		counts[i+1] += counts[i]
	}

	aux := make([]Student, l)
	for _, student := range students {
		aux[counts[student.Group]] = student
		counts[student.Group]++
	}

	// 转换为字符串数组的稳定排序结果
	res := make([]string, l)
	for _, stu := range aux {
		res = append(res, fmt.Sprintf("%d %d", stu.ID, stu.Group))
	}

	return res
}

func transforStudent(arr []string) []Student {
	if len(arr) == 0 {
		return nil
	}

	students := make([]Student, len(arr))
	for i := range len(arr) {
		record := arr[i]
		recordArr := strings.Split(record, " ")
		id, _ := strconv.Atoi(recordArr[0])
		group, _ := strconv.Atoi(recordArr[1])
		students = append(students, Student{
			ID:    id,
			Group: group,
		})
	}

	return students
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		strArr := make([]string, 0, n)
		count := 0
		for count < n {
			if !scanner.Scan() {
				break
			}
			record := strings.TrimSpace(scanner.Text())
			strArr = append(strArr, record)
			count++
		}
		students := transforStudent(strArr)
		res := keyIndexSort(students)
		for _, student := range res {
			fmt.Println(student)
		}
	}
}
