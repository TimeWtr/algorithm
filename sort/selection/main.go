package main

import "fmt"

// 选择排序，O(n2)的时间复杂度，空间复杂度O(1)
func sorted(arr []int, asc bool) []int {
	for i := 0; i < len(arr); i++ {
		targetIdx := i
		for j := i + 1; j < len(arr); j++ {
			if asc {
				if arr[targetIdx] > arr[j] {
					targetIdx = j
				}
			} else {
				if arr[targetIdx] < arr[j] {
					targetIdx = j
				}
			}
		}

		if targetIdx != i {
			arr[i], arr[targetIdx] = arr[targetIdx], arr[i]
		}
	}

	return arr
}

func main() {
	arr := []int{1, 101, 242, 11, 22, 1232, 78, 10, 188, 2, 2, 5}
	asc := sorted(arr, true)
	fmt.Println(asc)
	esc := sorted(arr, false)
	fmt.Println(esc)
}
