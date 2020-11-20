package sorting

import (
	"AlgorithmsByGo/utils"
	"testing"
)

var arr, _ = utils.Generate(10000,1, 1000000)

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++{
		_ = BubbleSort(arr)
	}
}


func BenchmarkInsertSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = InsertSort(arr)
	}
}

func BenchmarkSelectSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SelectSort(arr)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MergeSort(arr)
	}
}