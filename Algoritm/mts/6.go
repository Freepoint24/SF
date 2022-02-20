package main

import (
	"fmt"
	"sort"
)

//func main() {
//	value := "АаааИИИИцццЙа"
//	//сортировка входящей стркои
//	arr := []rune(value)
//	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
//
//	value2 := string(arr)
//	enc := Counter(value2)
//
//	fmt.Println(value)
//	fmt.Println(value2)
//	fmt.Println(enc)
//
//}

func Counter(s string) string {
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
