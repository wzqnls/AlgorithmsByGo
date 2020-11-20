package sorting

func QuickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(arr, l, r)
	QuickSort(arr, l, p-1)
	QuickSort(arr, p+1, r)

}

// patition return the value index of arr[l]
func partition(arr []int, l, r int) int {
	head := arr[l]
	index := l
	for i := l + 1; i <= r; i++ {
		if head > arr[i] {
			index++
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
	arr[l], arr[index] = arr[index], arr[l]
	return index
}
