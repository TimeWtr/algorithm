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

func handler(length, m int, arr []int) []int {
	head := buildCycleListNode(arr)
	if head == nil {
		return nil
	}

	var res []int
	count, removed := 0, 0
	prev, curr := head, head
	for removed < length {
		count++
		if count == m {
			res = append(res, curr.Val)
			count = 0
			m = curr.Val
			removed++
			prev.Next = curr.Next
			curr = prev.Next
		} else {
			prev = curr
			curr = curr.Next
		}
	}

	return res
}

func buildCycleListNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}
	curr := head
	for i := 1; i < len(arr); i++ {
		curr.Next = &ListNode{Val: arr[i]}
		curr = curr.Next
	}

	curr.Next = head
	return head
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		length, _ := strconv.Atoi(scanner.Text())

		if !scanner.Scan() {
			break
		}
		var arr []int
		for _, val := range strings.Split(strings.Trim(scanner.Text(), " "), ",") {
			v, _ := strconv.Atoi(val)
			arr = append(arr, v)
		}

		if !scanner.Scan() {
			break
		}
		m, _ := strconv.Atoi(scanner.Text())

		res := handler(length, m, arr)
		fmt.Println(res)
	}
}
