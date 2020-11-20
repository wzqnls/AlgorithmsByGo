package sorting

func MergeSort(arr []int) []int{
	mid := len(arr)/2
	if mid > 0 {
		left := MergeSort(arr[:mid])
		right := MergeSort(arr[mid:])
		return merge(left, right)
	} else {
		return arr
	}

}

func merge(left, right []int) []int{
	l, r := 0, 0
	var arr []int
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			arr = append(arr, left[l])
			l++
		} else {
			// left[l] >= right[r]
			arr = append(arr, right[r])
			r++
		}
	}
	if l < len(left) {
		arr = append(arr, left[l:]...)
	}
	if r < len(right) {
		arr = append(arr, right[r:]...)
	}
	return arr
}