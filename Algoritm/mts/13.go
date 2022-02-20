package main

import (
	"fmt"
	"sort"
)

func main() {
	value := "ЯЯЯБББддд"
	enc := Counter(value)
	fmt.Printf("%s\n%s\n", value, enc)
}

func Counter(s string) string {
	arr := []rune(s)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	s = string(arr)

	var zipp rune
	zippLen := 0
	res := ""
	for _, v := range s {
		if v != zipp {
			if zippLen > 0 {
				res += fmt.Sprintf("%c%d", zipp, zippLen)
			}
			zipp = v
			zippLen = 1
		} else {
			zippLen++
		}
	}
	res += fmt.Sprintf("%c%d", zipp, zippLen)
	return res
}
