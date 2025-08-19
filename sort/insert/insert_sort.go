package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)
	// 从第二个元素开始遍历（索引1到n-1）
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1 // 已排序序列的最后一个元素索引

		// 将大于key的元素向后移动一位
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key // 插入key到正确位置
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6}
	fmt.Println("排序前:", arr)
	insertionSort(arr)
	fmt.Println("排序后:", arr)
}
