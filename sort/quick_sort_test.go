package sort

import "testing"

func TestQuickSort(t *testing.T) {
	var arr = []int{5, 7, 3, 2, 7, 55, 33, 55, 13, 56, 758}
	QuickSort(arr)
	t.Log(arr)
	for k := range arr {
		if k == 0 {
			continue
		}
		if arr[k-1] > arr[k] {
			t.Error(arr, "排序失败")
			break
		}
	}
}
