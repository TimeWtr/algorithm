// 题目：环形链表
// 给定一个链表，返回链表开始入环的第一个节点。 从链表的头节点开始沿着 next 指针进入环的
// 第一个节点为环的入口节点。如果链表无环，则返回 null。
// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
// 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递
// 到函数中。

// 说明：不允许修改给定的链表。
//
// 示例1:
// 输入：head = [3,2,0,-4], pos = 1
// 输出：返回索引为 1 的链表节点
// 解释：链表中有一个环，其尾部连接到第二个节点。
//
// 示例2:
// 输入：head = [1,2], pos = 0
// 输出：返回索引为 0 的链表节点
// 解释：链表中有一个环，其尾部连接到第一个节点。
//
// 示例3:
// 输入：head = [1], pos = -1
// 输出：null
// 解释：链表中没有环。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getLinkedList(line string, pos int) *ListNode {
	strArr := strings.Split(strings.TrimRight(strings.TrimLeft(line, "["), "]"), ",")
	head := &ListNode{Val: -1}
	curr := head
	for i := range strArr {
		val, _ := strconv.Atoi(strArr[i])
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}

	next := head.Next
	if pos != -1 {
		for i := 0; i < pos; i++ {
			next = next.Next
		}
		curr.Next = next
	}

	return head.Next
}

// 方式1: 使用hashTable来检测是否有环
func checkIsCycleWithHashTable(head *ListNode) *ListNode {
	hashTable := map[*ListNode]struct{}{}
	curr := head
	for curr != nil {
		if curr.Next == nil {
			return nil
		}

		if _, ok := hashTable[curr]; ok {
			return curr
		}

		hashTable[curr] = struct{}{}
		curr = curr.Next
	}

	return nil
}

// 使用快慢双指针来判断是否有环
func checkIsCycleDoublePtr(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		if fast.Next == nil {
			return nil
		}

		// 慢指针走一步，快指针走两步
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// 出现了环
			ptr := head
			for ptr != slow {
				ptr = ptr.Next
				slow = slow.Next
			}
			return ptr
		}
	}

	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())

		if !scanner.Scan() {
			break
		}
		pos, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		head := getLinkedList(line, pos)
		ptr := checkIsCycleWithHashTable(head)
		fmt.Println(ptr)
		ptr1 := checkIsCycleDoublePtr(head)
		fmt.Println(ptr1)
	}
}
