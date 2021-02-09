package main

import (
	"fmt"
	"strconv"
)

func hashuint64(val uint64) uint64 {
	return (val % 1000)
}

func main() {
	x := "abcdERAAAAAAAAAAAAAAAAAAAAAAAAAA" // строка для преобразования
	fmt.Println("source string: ", x)

	var s string
	for _, b := range x {
		i := strconv.Itoa(int(b))
		s += i
	}
	sn, _ := strconv.ParseUint(s, 10, 64)
	fmt.Println("source number:", sn)

	res := hashuint64(sn)
	fmt.Println("result: ", res)
}
