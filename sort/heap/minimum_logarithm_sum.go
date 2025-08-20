package heap

import (
	"container/heap"
)

// sorted 使用选择排序实现对数组的升序和降序的排序
func sorted(arr []int, asc bool) []int {
	if len(arr) == 0 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		targetIdx := i
		for j := i + 1; j < len(arr); j++ {
			if asc {
				// 升序
				if arr[j] < arr[targetIdx] {
					targetIdx = j
				}
			} else {
				if arr[j] > arr[targetIdx] {
					targetIdx = j
				}
			}
		}

		if targetIdx != i {
			// 原地交换数据
			arr[i], arr[targetIdx] = arr[targetIdx], arr[i]
		}
	}

	return arr
}

type Pair struct {
	// 对数和
	sum int
	// 第一个数组下标
	i int
	// 第二个数组下标
	j int
}

type MinHeap []Pair

func (m MinHeap) Len() int            { return len(m) }
func (m MinHeap) Less(i, j int) bool  { return m[i].sum < m[j].sum }
func (m MinHeap) Swap(i, j int)       { m[i], m[j] = m[j], m[i] }
func (m *MinHeap) Push(x interface{}) { *m = append(*m, x.(Pair)) }
func (m *MinHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

func kSmallestPairs(array1, array2 []int, k int) int {
	m, n := len(array1), len(array2)
	// 边界条件
	if m <= 0 || n <= 0 || k <= 0 {
		return 0
	}

	// 初始化小顶堆
	h := MinHeap{}
	heap.Init(&h)

	// 将前K个对数写入到小顶堆中，这里注意两个数组都是升序的，所以可以将array2只取第一个下标的元素，
	// array1的下标顺序读取。
	for i := 0; i < m && i < k; i++ {
		heap.Push(&h, Pair{
			sum: array1[i] + array2[0],
			i:   i,
			j:   0,
		})
	}

	totalSum, count := 0, 0
	// 循环去除前K个对数
	for h.Len() > 0 && count < k {
		p := heap.Pop(&h).(Pair)
		totalSum += p.sum
		count++

		// 兜底实现，需要另外计算一步，array2下一个元素是否对数和是否比当前的对数和小
		if p.j+1 < n && count < k {
			heap.Push(&h, Pair{
				sum: array1[p.i] + array2[p.j+1],
				i:   p.i,
				j:   p.j + 1,
			})
		}
	}

	return totalSum
}
