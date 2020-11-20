package sorting

func InsertSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := i
		for j > 0 && tmp < arr[j-1] {
			arr[j] = arr[j-1]
			j -= 1
		}
		arr[j] = tmp
	}
	return arr
}

// InsertSort2 swap two values instead move value
func InsertSort2(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}

	}
	return arr
}
