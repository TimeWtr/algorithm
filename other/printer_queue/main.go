// //题目描述
// 有5台打印机打印文件，每台打印机有自己的待打印队列。
// 因为打印的文件内容有轻重缓急之分，所以队列中的文件有1~10不同的代先级，其中数字越大优先级越高。
// 打印机会从自己的待打印队列中选择优先级最高的文件来打印。
// 如果存在两个优先级一样的文件，则选择最早进入队列的那个文件。
// 现在请你来模拟这5台打印机的打印过程。

// 输入描述
// 每个输入包含1个测试用例，
// 每个测试用例第一行给出发生事件的数量N（0 < N < 1000）。

// 接下来有 N 行，分别表示发生的事件。共有如下两种事件：

// “IN P NUM”，表示有一个拥有优先级 NUM 的文件放到了打印机 P 的待打印队列中。（0< P <= 5, 0 < NUM <= 10)；
// “OUT P”，表示打印机 P 进行了一次文件打印，同时该文件从待打印队列中取出。（0 < P <= 5）。
// 输出描述
// 对于每个测试用例，每次”OUT P”事件，请在一行中输出文件的编号。
// 如果此时没有文件可以打印，请输出”NULL“。
// 文件的编号定义为”IN P NUM”事件发生第 x 次，此处待打印文件的编号为x。编号从1开始。
//
// 用例1
// 输入
// 7
// IN 1 1
// IN 1 2
// IN 1 3
// IN 2 1
// OUT 1
// OUT 2
// OUT 2
// 输出
// 3
// 4
// NULL
// 用例2
// 输入
// 5
// IN 1 1
// IN 1 3
// IN 1 1
// IN 1 3
// OUT 1
// 输出
// 2
//
// 解题思路：
// 使用优先级队列（大顶堆）来实现单个打印机，打印机的排序条件：
// 1. 高优先级优先出队
// 2. 每一个入队的文件都有一个文件序列号，序列号低的优先出队，即文件的入队顺序

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Sequence int64

type PrinterJob struct {
	// 文件序列号
	Sequence int64
	// 文件优先级
	Priority int
}

type Printer []*PrinterJob

func (p Printer) Len() int { return len(p) }
func (p Printer) Less(i, j int) bool {
	// 优先选择优先级更高的
	if p[i].Priority != p[j].Priority {
		return p[i].Priority > p[j].Priority
	}

	// 优先级相同的情况下序列号小的优先
	return p[i].Sequence < p[j].Sequence
}
func (p Printer) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p *Printer) Push(v any) {
	*p = append(*p, v.(*PrinterJob))
}
func (p *Printer) Pop() any {
	old := *p
	l := len(old)
	v := old[l-1]
	*p = old[:l-1]
	return v
}

var printerQueue []*Printer

func initQueue() {
	printerQueue = make([]*Printer, 5)
	for i := range printerQueue {
		printerQueue[i] = &Printer{}
		heap.Init(printerQueue[i])
	}
}

func handler(event string) {
	fields := strings.Fields(event)
	number, _ := strconv.Atoi(fields[1])
	printer := printerQueue[number-1]
	if fields[0] == "OUT" {
		// 打印
		if printer.Len() == 0 {
			fmt.Println("NULL")
		} else {
			job := heap.Pop(printer).(*PrinterJob)
			fmt.Println(job.Sequence)
		}

		return
	}

	// 写入
	priority, _ := strconv.Atoi(fields[2])
	Sequence++
	heap.Push(printer, &PrinterJob{
		Sequence: Sequence,
		Priority: priority,
	})
}

func main() {
	initQueue()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		num, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		count := 0
		for count < num {
			if !scanner.Scan() {
				break
			}

			event := strings.TrimSpace(scanner.Text())
			handler(event)
			count++
		}
		Sequence = 0
	}
}
