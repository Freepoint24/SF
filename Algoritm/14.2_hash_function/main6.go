package main

import (
	"fmt"
)

func hashuint64(val uint64) uint64 {
	return (val % 1000)
}

func main() {
	var err error
	var val uint64
	//ввод числа
	fmt.Print("Введите число ")
	fmt.Scan(&val)
	_, err = fmt.Scanln(&val)
	if err != nil {
		fmt.Printf("ошибка при вводе числа")
		return
		//x := "abcdERAAAAAAAAAAAAAAAAAAAAAAAAAA" // строка для преобразования
		//fmt.Println("source string: ", x)
		//var s string
		//for _, b := range x {
		//	i := strconv.Itoa(int(b))
		//	s += i
		//}
		//sn, _ := strconv.ParseUint(s, 10, 64)
		//fmt.Println("source number:", sn)
		res := hashuint64(val)
		fmt.Println("result: ", res)
	}
}
