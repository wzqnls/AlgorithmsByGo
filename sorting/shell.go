package sorting


func ShellSort(arr []int) []int {
	for step := len(arr) / 2; step > 0 ; step /= 2 {
		for i := step; i < len(arr); i++ {
			for j := i; j >= step && arr[j] < arr[j - step]; j -= step {
				arr[j], arr[j - step] = arr[j - step], arr[j]
			}
		}
	}
	return arr
}
