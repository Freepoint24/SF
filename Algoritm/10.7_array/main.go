package main

import (
	"fmt"
	"math"
)

func main() {

	arr := []int{34, 67, 12, 3, 4, 9, 108, 6, 4, 3, 1, 100, 39}
	maxVal := findMax(arr)
	fmt.Printf("Max value in array %v -> %v \n", arr, maxVal)

	// Max value in array [34 67 12 3 4 9 108 6 4 3 1 100 39] -> 108
}

func findMax(array []int) (max int) {
	max = math.MinInt64
	for _, val := range array {
		if val > max {
			max = val
		}
	}

	return max
}
