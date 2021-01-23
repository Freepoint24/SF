package main

import (
	"fmt"
	//импортитуется пакет smplStr из модуля kvadratStr
	"kvadratStr/smplStr"
)

func main() {
	var storona1, storona2 float64
	var err error

	//ввод числа 1, если введено не число float 64,
	//то возвращается перемнная err и выводится сообщине "ошибка при вводе первого числа"
	fmt.Print("Введите число 1:")
	_, err = fmt.Scanln(&storona1)
	if err != nil {
		fmt.Printf("ошибка при вводе первого числа")
		return
	}
	//ввод числа 2. если введено не число float 64,
	//то возвращается перемнная err и выводится сообщине "ошибка при вводе первого числа"
	fmt.Print("Введите число 2:")
	_, err = fmt.Scanln(&storona2)
	if err != nil {
		fmt.Printf("ошибка при вводе первого числа")
		return
	}

	//это и есть конструктор?
	// Переменная Kvadratt возвращает экземляр структур NewKvadrat из пакета smplStr
	// Переменная resultat получает доступ к методу Kvadr с перемнными storona1 и storona2 который возвращает их умножение
	// метод Println из пакета fmt выводит значиение переменной resultat
	Kvadratt := smplStr.NewKvadrat()
	resultat := Kvadratt.Kvadr(storona1, storona2)
	fmt.Println("Результат:", resultat)

}
