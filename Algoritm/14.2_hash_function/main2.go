package main

import (
	"fmt"
)

func main() {
	var number1 uint64

	var err error

	//ввод числа 1
	fmt.Print("Введите число:")
	//fmt.Scan(&number1)
	_, err = fmt.Scanln(&number1)
	if err != nil {
		fmt.Printf("ошибка при вводе первого числа")
		return

		// блок вычислений
		val = number1 % 1000
		fmt.Println(val)
		//return val
	}
}
