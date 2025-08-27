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

// 文件序列编号
var Sequence int
var queues []*Printer

// initPrinters 初始化打印机
func initPrinters(num int) {
	queues = make([]*Printer, num)
	for i := range queues {
		queues[i] = &Printer{}
		heap.Init(queues[i])
	}
}

type PrintJob struct {
	// 优先级
	Priority int
	// 文件序列号
	Sequence int
}

type Printer []*PrintJob

func (p Printer) Len() int { return len(p) }
func (p Printer) Less(i, j int) bool {
	// 先比较优先级
	if p[i].Priority != p[j].Priority {
		return p[i].Priority > p[j].Priority
	}

	// 相同的优先级比较序列号
	return p[i].Sequence < p[j].Sequence
}
func (p Printer) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p *Printer) Push(v any) {
	*p = append(*p, v.(*PrintJob))
}
func (p *Printer) Pop() any {
	old := *p
	l := len(old)
	v := old[l-1]
	*p = old[:l-1]
	return v
}

func handler(et string, printer int, priority ...int) {
	p := queues[printer-1]
	if et == "IN" {
		// 输入
		Sequence++
		job := PrintJob{
			Priority: priority[0],
			Sequence: Sequence,
		}
		heap.Push(p, &job)
	} else if et == "OUT" {
		// 输出
		if p.Len() == 0 {
			fmt.Printf("NULL")
			return
		}
		v := heap.Pop(p).(*PrintJob)
		fmt.Println(v.Sequence)
	}
}

func main() {
	// 初始化5台打印机和堆
	initPrinters(5)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// 接收事件数量
		if !scanner.Scan() {
			break
		}
		num, _ := strconv.Atoi(scanner.Text())

		// 接收事件
		count := 0
		var (
			eventType string
			printer   int
			priority  int
		)
		for count < num {
			if !scanner.Scan() {
				break
			}

			eventArr := strings.Split(strings.TrimSpace(scanner.Text()), " ")
			eventType = eventArr[0]
			printer, _ = strconv.Atoi(eventArr[1])
			if len(eventArr) == 3 {
				priority, _ = strconv.Atoi(eventArr[2])
			}
			handler(eventType, printer, priority)
			count++
		}
	}
}
