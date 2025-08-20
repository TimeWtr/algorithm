package bubble

func SortBubble(arr []int, asc bool) []int {
	if len(arr) <= 1 {
		return arr
	}

	for i := 0; i < len(arr)-1; i++ {
		swapped := false
		for j := 0; j < len(arr)-i-1; j++ {
			if asc {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					swapped = true
				}
			} else {
				if arr[j] < arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					swapped = true
				}
			}
		}

		if !swapped {
			break
		}
	}

	return arr
}
