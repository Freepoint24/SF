package main

import (
	"fmt"
)

func main() {
	firstarray := []int{1, 9, 12, 16, 13, 11, 14, 18, 8, 12}
	secondarray := []int{9, 7, 15, 8, 14, 19, 34, 18, 13, 11}
	var finalarray []int

	for _, a1 := range firstarray {
		for _, a2 := range secondarray {
			if a1 == a2 {
				finalarray = append(finalarray, a2)
			}
		}
	}

	fmt.Println("matching no is", finalarray)

}
