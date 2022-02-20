package main

import (
	"fmt"
	"sort"
)

func main() {

	value := "Уууууцйщллщлввсв"
	arr := []rune(value)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	fmt.Println(value)
	value2 := string(arr)
	enc := encode(value2)
	fmt.Printf("%s\n%s\n", value2, enc)
}

func encode(s string) string {
	var curr rune
	currLen := 0
	res := ""
	for _, v := range s {
		if v != curr {
			if currLen > 0 {
				res += fmt.Sprintf("%c%d", curr, currLen)
			}
			curr = v
			currLen = 1
		} else {
			currLen++
		}
	}
	res += fmt.Sprintf("%c%d", curr, currLen)
	return res
}
