package main

import (
	"fmt"
	"strconv"
)

func hashstr(val uint64) uint64 {
	return (val % 1000)
}


type hashmap struct {
	key string
	val string
	ok bool
}

func (h *hashmap) Set(key, val string) {

}

func (h *hashmap) Get(key string) (value string, ok bool) {
	return 
}

func (h *hashmap) Delete(key string) {
}



func main() {
	h := "abc" // строка для преобразования
	fmt.Println("source string: ", h)

	var s string
	for _, b := range h {
		i := strconv.Itoa(int(b))
		s += i
	}
	sn, _ := strconv.ParseUint(s, 10, 64)
	fmt.Println("source number:", sn) //входные данные

	res := hashstr(sn)
	fmt.Println("result: ", res) // хэш

}


//tools
