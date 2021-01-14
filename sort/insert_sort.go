package sort

func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

//递归版 二分查找
func InsertSort2(arr []int) {
	if len(arr) > 2 {
		InsertSort2(arr[:len(arr)-1])
	}

	// 利用二分查找，在待排元素左侧找到合适的插入位置
	p := suitableIndex2(arr)

	// 如果最合适的位置不是待排元素当前位置，那就一次把合适位置后的元素都向后移动一位
	if p != len(arr)-1 {
		temp := arr[len(arr)-1]
		for j := len(arr) - 1; j > p; j-- {
			arr[j] = arr[j-1]
		}
		arr[p] = temp
	}

}

//二分查找递归版
func suitableIndex(list []int, start int, end int, current int) int {
	// 比到最后没的比的时候就去对比下当前位置与待排元素的大小，并返回较大值的位置
	if start >= end {
		if list[start] < list[current] {
			return current //比所有的元素都大
		} else {
			return start //比所有的元素都小
		}
	}

	center := (end-start)/2 + start

	// 如果中间的元素比较大，就继续向左侧寻找。反之则向右
	if list[center] > list[current] {
		return suitableIndex(list, start, center, current)
	} else {
		return suitableIndex(list, center+1, end, current)
	}

}

//二分查找循环版
func suitableIndex2(list []int) int {
	key := len(list) - 1
	start, end := 0, len(list)-2
	for start < end {
		center := (end-start)/2 + start
		if list[key] < list[center] {
			end = center
		} else {
			start = center + 1
		}
	}
	if list[key] > list[start] {
		return key
	} else {
		return start
	}
}
