package sort

func merge(arr []int, left, mid, right int) {
	L := make([]int, mid-left+1)
	R := make([]int, right-mid)
	copy(L, arr[left:mid+1])
	copy(R, arr[mid+1:right+1])
	//遇到边界问题，需要结合实例来逐步分析
	//for i := 0; i <= mid-left; i++ {
	//	L = append(L, arr[left+i])
	//}
	//for i := 0; i < right-mid; i++ {
	//	R = append(R, arr[mid+i+1])
	//}
	//L = append(L, math.MaxInt32)
	//R = append(R, math.MaxInt32)
	i, j := 0, 0

	for k := left; k <= right; k++ {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		if i >= len(L) {
			copy(arr[k+1:right+1], R[j:])
			break
		}
		if j >= len(R) {
			copy(arr[k+1:right+1], L[i:])
			break
		}
	}

}
func MergeSort(arr []int, left, right int) {
	if left < right {
		mid := (left + right) >> 1
		MergeSort(arr, left, mid)
		MergeSort(arr, mid+1, right)
		merge(arr, left, mid, right)
	}
}
