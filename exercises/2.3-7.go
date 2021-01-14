package exercises

import (
	"arithmetic/sort"
)

func SearchSum(arr []int, sum int) bool {
	sort.MergeSort(arr, 0, len(arr)-1) //归并排序 O(n lgn) 复杂度
	if arr[0] > sum {
		return false
	}
	for i := 0; i < len(arr)-1; i++ { //n次循环 加二分查找 也是 O(n lgn) 复杂度
		ret := binarySearch(arr[i+1:], sum-arr[i])
		if ret != -1 {
			return true
		}
	}
	return false
}
func binarySearch(arr []int, key int) int {
	start, end := 0, len(arr)-1
	for start <= end {
		mid := start + (end-start)/2
		if key > arr[mid] {
			start = mid + 1
		} else if key < arr[mid] {
			end = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
