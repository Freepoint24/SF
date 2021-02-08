package main

import (
	"fmt"
)

func main() {
	//	var err error
	////	var z int64
	//	//ввод числа
	//	fmt.Print("Введите число ")
	//	fmt.Scan(&z)
	//	_, err = fmt.Scanln(&z)
	//	if err != nil {
	//		fmt.Printf("ошибка при вводе первого числа")
	//		return
	//	}

}

func hashint64(val int64) int64 {
	z := 164654
	x := 1000
	val := z % x
	fmt.Println(val)
	return val
}
