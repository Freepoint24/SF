package main

import (
	"fmt"
)

func main() {
	type void struct{}
	var member void
	firstarray := map[int]void{1: member, 9: member, 12: member, 16: member, 13: member, 11: member, 14: member, 18: member, 8: member}
	secondarray := map[int]void{9: member, 7: member, 15: member, 8: member, 14: member, 19: member, 34: member, 18: member, 13: member, 11: member}

	var finalarray []int
	for k := range firstarray {
		if _, ok := secondarray[k]; ok {
			finalarray = append(finalarray, k)
		}
	}

	fmt.Println("matching no is", finalarray)
}
