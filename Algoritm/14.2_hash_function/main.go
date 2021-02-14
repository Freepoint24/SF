package main

import (
	"fmt"
)

func main() {
	var err error
	var z int
	//ввод числа
	fmt.Print("Введите число ")
	fmt.Scan(&z)
	_, err = fmt.Scanln(&z)
	if err != nil {
		fmt.Printf("ошибка при вводе числа")
		return
	}

}

//func hashint64(val int) uint {
	//z := 164654
	//x := 1000
	val = z % 1000 //x
	fmt.Println(val)
	return uint(val)

}
