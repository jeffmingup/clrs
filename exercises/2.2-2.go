package exercises

func SelectSort(arr []int) {
	for k := range arr {
		min := k
		for i := k; i < len(arr); i++ {
			if arr[i] < arr[min] {
				min = i
			}
		}
		arr[k], arr[min] = arr[min], arr[k]
	}
}
