package main

import (
	"fmt"
	//"math"
)

func main() {

	arr := []int{34, 67, 12, 3, 4, 9, 108, 6, 4, 3, 1, 100, 39}
	maxVal := findMax(arr)
	fmt.Printf("Max value in array %v -> %v \n", arr, maxVal)

	// Max value in array [34 67 12 3 4 9 108 6 4 3 1 100 39] -> 108
}

func findMax(array []int) (max int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found max in empty slice")
	}

	max = array[0]
	for _, val := range array[1:] {
		if val > max {
			max = val
		}
	}

	return max, nil
}