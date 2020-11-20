package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func Generate(length, min, max int) ([]int, error) {
	if length <= 0 {
		return []int{}, fmt.Errorf("error length, must bigger than 0")
	}
	if min >= max {
		return []int{}, fmt.Errorf("error scope, min must less than max value")
	}
	var arr []int
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		value := rand.Intn(max-min) + min
		arr = append(arr, value)
	}
	return arr, nil
}
