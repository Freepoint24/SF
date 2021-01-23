package main

import (
	"fmt"
	"kvadrat/SmplStr"
)

func main() {
	var storona1, storona2 float64
	var err error

	//ввод числа 1
	fmt.Print("Введите число 1:")
	_, err = fmt.Scanln(&storona1)
	if err != nil {
		fmt.Printf("ошибка при вводе первого числа")
		return
	}
	//ввод числа 2
	fmt.Print("Введите число 2:")
	_, err = fmt.Scanln(&storona2)
	if err != nil {
		fmt.Printf("ошибка при вводе первого числа")
		return
	}

	Kvadrat := smplStr.NewKvadrat()
	result := Kvadrat(storona1, storona2)
	fmt.Println(result)

}
