package main

import (
	"AlgorithmsByGo/sorting"
	"AlgorithmsByGo/utils"
	"fmt"
)

func main() {
	arr, _ := utils.Generate(10, 1, 100)
	//fmt.Println(sorting.BubbleSort(arr))
	//fmt.Println(sorting.SelectSort(arr))
	//fmt.Println(sorting.InsertSort(arr))
	//fmt.Println(sorting.MergeSort(arr))
	//fmt.Println(sorting.ShellSort(arr))

	sorting.QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
