package main

import (
	"fmt"
	"kvadratStr/smplStr"
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

	//это и есть конструктор?
	// Kvadrat возвращает экземляр структур NewKvadrat из пакета smplStr
	//??? resultat получает доступ

	Kvadrat := smplStr.NewKvadrat()
	resultat := Kvadrat.Kvadr(storona1, storona2)
	fmt.Println("Результат:", resultat)

}
